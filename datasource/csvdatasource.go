package datasource

import (
	"bmc-test-golang-service/model"
	"bmc-test-golang-service/util"
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"sync"
)

type CSVDataSource struct {
	FilePath      string
	Data          []model.PassengerInfo
	IsCacheLoaded bool
	mu            sync.RWMutex // Mutex for concurrent access to the data slice and cache flag.
}

func (ds *CSVDataSource) ReadAllData() ([]model.PassengerInfo, error) {
	ds.mu.RLock()
	if ds.IsCacheLoaded {
		// Return a copy of the cached data.
		dataCopy := make([]model.PassengerInfo, len(ds.Data))
		copy(dataCopy, ds.Data)
		ds.mu.RUnlock()
		return dataCopy, nil
	}
	ds.mu.RUnlock()

	file, err := os.Open(ds.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	var data []model.PassengerInfo

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		passenger := model.PassengerInfo{
			PassengerID: record[0],
			Survived:    util.ParseInt(record[1]),
			Pclass:      util.ParseInt(record[2]),
			Name:        record[3],
			Sex:         record[4],
			Age:         util.ParseFloat32(record[5]),
			SibSb:       util.ParseInt(record[6]),
			Parch:       util.ParseInt(record[7]),
			Ticket:      record[8],
			Fare:        util.ParseFloat64(record[9]),
			Cabin:       record[10],
			Embarked:    record[11],
		}

		data = append(data, passenger)
	}

	ds.mu.Lock()
	ds.Data = data
	ds.IsCacheLoaded = true
	ds.mu.Unlock()

	dataCopy := make([]model.PassengerInfo, len(data))
	copy(dataCopy, data)

	return dataCopy, nil
}

func (ds *CSVDataSource) ReadAllDataPageable(c *gin.Context) ([]model.PassengerInfo, error) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter"})
		return nil, err
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pageSize parameter"})
		return nil, err
	}

	allData, err := ds.ReadAllData()
	if err != nil {
		return nil, err
	}

	startIndex := (pageInt - 1) * pageSizeInt
	endIndex := startIndex + pageSizeInt

	if startIndex >= len(allData) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Page out of bounds"})
		return nil, err
	}

	var pageData []model.PassengerInfo
	if endIndex > len(allData) {
		pageData = allData[startIndex:]
	} else {
		pageData = allData[startIndex:endIndex]
	}

	return pageData, err
}

func (ds *CSVDataSource) ReadDataByID(passengerID string) (*model.PassengerInfo, error) {
	data, err := ds.ReadAllData()
	if err != nil {
		return nil, err
	}

	for _, passenger := range data {
		if passenger.PassengerID == passengerID {
			return &passenger, nil
		}
	}

	return nil, fmt.Errorf("Passenger not found with ID: %d", passengerID)
}

func (ds *CSVDataSource) ReadAllFares() ([]float64, error) {
	data, err := ds.ReadAllData()
	if err != nil {
		return nil, err
	}

	var fares []float64
	for _, info := range data {
		fares = append(fares, info.Fare)
	}

	sort.Float64s(fares)
	return fares, err
}

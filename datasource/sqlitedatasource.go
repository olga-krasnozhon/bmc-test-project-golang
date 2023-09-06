package datasource

import (
	"bmc-test-golang-service/model"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type SQLiteDataSource struct {
	DB *gorm.DB
}

func (ds *SQLiteDataSource) ReadAllDataPageable(c *gin.Context) ([]model.PassengerInfo, error) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt <= 0 {
		return nil, errors.New("Invalid page parameter")
	}

	pageSizeInt, err := strconv.Atoi(pageSize)
	if err != nil || pageSizeInt <= 0 {
		return nil, errors.New("Invalid pageSize parameter")
	}

	offset := (pageInt - 1) * pageSizeInt

	var passengers []model.PassengerInfo
	if err := ds.DB.Offset(offset).Limit(pageSizeInt).Find(&passengers).Error; err != nil {
		return nil, err
	}

	return passengers, nil
}

func (ds *SQLiteDataSource) ReadDataByID(passengerID string) (*model.PassengerInfo, error) {
	var passenger model.PassengerInfo

	if err := ds.DB.First(&passenger, passengerID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Passenger not found")
		}
		return nil, err
	}

	return &passenger, nil
}

func (ds *SQLiteDataSource) ReadAllFares() ([]float64, error) {
	var sortedFares []float64

	query := "SELECT fare FROM passenger_infos ORDER BY fare"
	if err := ds.DB.Raw(query).Pluck("fare", &sortedFares).Error; err != nil {
		return nil, err
	}

	return sortedFares, nil
}

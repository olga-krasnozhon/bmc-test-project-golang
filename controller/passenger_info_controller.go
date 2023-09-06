package controller

import (
	"bmc-test-golang-service/datasource"
	"bmc-test-golang-service/model"
	"bmc-test-golang-service/util"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"net/http"
	"strings"
)

// GetPassengerInfo @Summary Get Passenger Info
// @Description Get Passenger Info by passenger ID
// @ID GetPassengerInfo
// @Produce json
// @Param passengerId path int true "Passenger ID"
// @Success 200 {object} model.PassengerInfo "Passenger Info"
// @Failure 404 {object} gin.H "Passenger not found"
// @Router /passengers/v1/info/{passengerId} [get]
func GetPassengerInfo(c *gin.Context, ds datasource.DataSource) {
	id := c.Param("passengerId")

	data, err := ds.ReadDataByID(id)
	if err != nil {
		c.JSON(404, gin.H{
			"Error": "Passenger not found: " + err.Error(),
		})
		return
	}
	c.JSON(200, data)
}

// GetAllPassengerInfo @Summary Get all passenger information
// @Description Get a list of all passenger information.
// @ID getAllPassengerInfo
// @Produce json
// @Success 200 {array} model.PassengerInfo
// @Router /passengers/v1/info [get]
func GetAllPassengerInfo(c *gin.Context, ds datasource.DataSource) {
	pageData, err := ds.ReadAllDataPageable(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve passenger information"})
		return
	}

	c.JSON(http.StatusOK, pageData)
}

// GetFaresHistogram @Summary Get fares histogram
// @Description Get a histogram of fare prices with percentiles.
// @ID getFaresHistogram
// @Produce text/html
// @Success 200 {string} string "html"
// @Router /passengers/v1/info/fares/histogram [get]
func GetFaresHistogram(c *gin.Context, ds datasource.DataSource) {

	farePrices, _ := ds.ReadAllFares()
	percentiles := []float64{5, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	var histogramData []model.HistogramData
	for _, p := range percentiles {
		idx := int((p / 100.0) * float64(len(farePrices)))
		count := idx + 1 // Count of items under the percentile
		histogramData = append(histogramData, model.HistogramData{Percentile: p, Count: count})
	}

	tmpl := util.CreateHistogramHtmlTemplate()

	labels := []string{}
	data := []int{}

	for _, item := range histogramData {
		labels = append(labels, fmt.Sprintf("%.1f%%", item.Percentile))
		data = append(data, item.Count)
	}

	dataMap := map[string]interface{}{
		"Labels": labels,
		"Data":   data,
	}

	t, _ := template.New("histogram").Parse(tmpl)
	t.Execute(c.Writer, dataMap)
}

// GetPassengerInfoByAttributes @Summary Get passenger information by attributes
// @Description Get passenger information filtered by specific attributes.
// @ID getPassengerInfoByAttributes
// @Produce json
// @Param passengerId query string true "Passenger ID"
// @Param attributes query string true "Comma-separated list of attribute names"
// @Success 200 {object} model.PassengerInfoDTO "Passenger information"
// @Failure 400 {object} gin.H "Bad request"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /passenger/attributes [get]
func GetPassengerInfoByAttributes(c *gin.Context, ds datasource.DataSource) {
	attributeNames := c.DefaultQuery("attributes", "")
	passengerID := c.DefaultQuery("passengerId", "")

	if attributeNames == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Please provide at least one attribute name"})
		return
	}

	attributeNamesSlice := strings.Split(attributeNames, ",")

	for i, name := range attributeNamesSlice {
		attributeNamesSlice[i] = strings.ToLower(name)
	}

	if !util.ValidateAttributes(attributeNamesSlice) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid attribute name(s)"})
		return
	}

	passenger, err := ds.ReadDataByID(passengerID)
	passengerDTO := util.CreatePassengerInfoUsingPassedAttributes(*passenger, attributeNamesSlice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving passenger information"})
		return
	}

	c.JSON(http.StatusOK, passengerDTO)
}

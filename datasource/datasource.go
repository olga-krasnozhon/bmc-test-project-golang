package datasource

import (
	"bmc-test-golang-service/model"
	"github.com/gin-gonic/gin"
)

type DataSource interface {
	ReadAllDataPageable(c *gin.Context) ([]model.PassengerInfo, error)
	ReadDataByID(passengerID string) (*model.PassengerInfo, error)
	ReadAllFares() ([]float64, error)
}

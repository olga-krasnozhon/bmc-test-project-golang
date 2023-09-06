package router

import (
	"bmc-test-golang-service/controller"
	"bmc-test-golang-service/datasource"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"sync"
)

func InitializeRouter(ds datasource.DataSource, maxWorkers int) *gin.Engine {
	r := gin.Default()

	workerPool := make(chan struct{}, maxWorkers)

	url := ginSwagger.URL("/swagger/swagger.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	executeInPool := func(handlerFunc func(*gin.Context, datasource.DataSource), c *gin.Context) {
		workerPool <- struct{}{}
		defer func() {
			<-workerPool
		}()
		handlerFunc(c, ds)
	}

	r.GET("/passengers/v1/info/:passengerId", func(c *gin.Context) {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			executeInPool(controller.GetPassengerInfo, c)
		}()
		wg.Wait()
	})

	r.GET("/passengers/v1/info", func(c *gin.Context) {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			executeInPool(controller.GetAllPassengerInfo, c)
		}()
		wg.Wait()
	})

	r.GET("/passengers/v1/info/:passengerId/:attributes", func(c *gin.Context) {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			executeInPool(controller.GetPassengerInfoByAttributes, c)
		}()
		wg.Wait()
	})

	r.GET("/passengers/v1/info/fare-histogram", func(c *gin.Context) {
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			executeInPool(controller.GetFaresHistogram, c)
		}()
		wg.Wait()
	})

	return r
}

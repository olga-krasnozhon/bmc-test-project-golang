package main

import (
	"bmc-test-golang-service/configuration"
	"bmc-test-golang-service/database"
	"bmc-test-golang-service/datasource"
	_ "bmc-test-golang-service/docs"
	"bmc-test-golang-service/model"
	_ "bmc-test-golang-service/model"
	"bmc-test-golang-service/server/router"
	"bmc-test-golang-service/util"
	"flag"
	"log"
	"runtime"
	_ "strconv"
)

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	var ds datasource.DataSource
	if config.UseCSV {
		ds = &datasource.CSVDataSource{
			FilePath:      "titanic.csv",
			Data:          make([]model.PassengerInfo, 0),
			IsCacheLoaded: false,
		}
	} else if config.UseSQLite {
		db := database.SetupDatabase()
		util.ImportCsvFileIntoDatabase(db)
		ds = &datasource.SQLiteDataSource{DB: db}
	} else {
		// TODO
	}

	maxWorkers := runtime.NumCPU()

	r := router.InitializeRouter(ds, maxWorkers)
	r.Run(":8080")
}

func loadConfig() (*configuration.Config, error) {
	configFile := flag.String("config", "", "Path to the configuration file (JSON)") // ./bmc-test-golang-service -config /Users/olya/work/bmc-test-golang-service/configuration/config.json
	flag.Parse()

	if *configFile == "" {
		log.Fatal("Please specify the path to the configuration file using the -config flag.")
	}

	config, err := configuration.LoadConfig(*configFile)
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	return config, err
}

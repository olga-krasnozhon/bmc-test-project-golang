package database

import (
	"bmc-test-golang-service/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func SetupDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("mydatabase.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	migrator := db.Migrator()
	if migrator.HasTable(&model.PassengerInfo{}) {
		if err := migrator.DropTable(&model.PassengerInfo{}); err != nil {
			log.Fatal(err)
		}
	}

	if err := db.AutoMigrate(&model.PassengerInfo{}); err != nil {
		log.Fatal(err)
	}

	return db
}

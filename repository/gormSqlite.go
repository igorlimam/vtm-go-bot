package repository

import (
	"log"
	"vtm-go-bot/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectGormSqlite() {
	db, err := gorm.Open(sqlite.Open("vtmgo.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}

func CheckDDL() {
	if DB == nil {
		ConnectGormSqlite()
	}
	DB.AutoMigrate(&model.Discipline{})
	DB.AutoMigrate(&model.Power{})
	DB.AutoMigrate(&model.Clan{})
}

func InsertIntoTable(tableInstance interface{}) {
	err := DB.Create(tableInstance)
	if err.Error != nil {
		log.Fatalf("Error inserting into table: %v", err.Error)
	}
}

func GetAll(tableInstance interface{}) {
	err := DB.Find(tableInstance)
	if err.Error != nil {
		log.Fatalf("Error retrieving all records: %v", err.Error)
	}
}

func GetByID(tableInstance interface{}, id uint) {
	err := DB.First(tableInstance, id)
	if err.Error != nil {
		log.Fatalf("Error retrieving record by ID: %v", err.Error)
	}
}

func GetByField(tableInstance interface{}, fieldName string, value interface{}) {
	err := DB.Where(fieldName+" = ?", value).Find(tableInstance)
	if err.Error != nil {
		log.Fatalf("Error retrieving records by field %s: %v", fieldName, err.Error)
	}
}

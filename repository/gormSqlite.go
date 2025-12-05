package repository

import (
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
}

func InsertIntoTable(tableInstance interface{}) {
	DB.Create(tableInstance)
}

/*
func getById(tableInstance interface{}, id uint) {
	db := ConnectGormSqlite()
	db.First(tableInstance, id)
}
*/

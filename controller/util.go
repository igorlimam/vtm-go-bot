package controller

import (
	"log"
	"vtm-go-bot/service"
)

func CheckDDLController() {
	service.CheckDDLService()
	log.Println("Database schema checked and updated if necessary.")
}

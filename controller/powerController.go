package controller

import (
	"log"
	"vtm-go-bot/model"
	"vtm-go-bot/service"

	"github.com/bwmarrin/discordgo"
)

func AddPower(session *discordgo.Session, interaction *discordgo.InteractionCreate, disciplineId string) string {
	log.Println("AddPower command invoked")
	status := service.AddPowerService(interaction, disciplineId)["status"]
	return status
}

func GetAllPowers() []model.Power {
	return service.GetAllPowers()
}

func GetDisciplinePowersByID(disciplineID string) []model.Power {
	return service.GetDisciplinePowersByID(disciplineID)
}

func GetPowerById(powerID string) model.Power {
	return service.GetPowerById(powerID)
}

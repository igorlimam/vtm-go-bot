package controller

import (
	"log"
	"vtm-go-bot/model"
	"vtm-go-bot/service"

	"github.com/bwmarrin/discordgo"
)

func AddPower(session *discordgo.Session, interaction *discordgo.InteractionCreate, disciplineId string) string {
	log.Println("AddPower command invoked")
	status := service.AddPowerService(interaction, disciplineId, "")["status"]
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

func UpdatePower(s *discordgo.Session, interaction *discordgo.InteractionCreate, powerID string, disciplineID string) string {
	log.Println("UpdatePower command invoked")
	status := service.AddPowerService(interaction, disciplineID, powerID)["status"]
	return status
}

func DeletePower(s *discordgo.Session, interaction *discordgo.InteractionCreate, powerID string) string {
	status := service.DeletePowerService(interaction, powerID)["status"]
	return status
}

package controller

import (
	"log"
	"vtm-go-bot/service"

	"github.com/bwmarrin/discordgo"
)

func AddPower(session *discordgo.Session, interaction *discordgo.InteractionCreate, disciplineId string) string {
	log.Println("AddPower command invoked")
	status := service.AddPowerService(interaction, disciplineId)["status"]
	return status
}

package controller

import (
	"log"
	"vtm-go-bot/service"

	"github.com/bwmarrin/discordgo"
)

func AddDiscipline(interaction *discordgo.InteractionCreate) map[string]string {
	log.Println("AddDiscipline command invoked")

	service.AddDisciplineService(interaction)

	return map[string]string{
		"status": "Discipline added successfully",
	}
}

package controller

import (
	"vtm-go-bot/service"

	"github.com/bwmarrin/discordgo"
)

func AddMerit(s *discordgo.Session, interaction *discordgo.InteractionCreate) string {
	status := service.AddMeritService(interaction, "")["status"]
	return status
}

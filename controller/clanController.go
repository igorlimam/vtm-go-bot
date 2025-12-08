package controller

import (
	"vtm-go-bot/service"

	"github.com/bwmarrin/discordgo"
)

func AddClan(s *discordgo.Session, interaction *discordgo.InteractionCreate, disciplineIDsSuffix string) {
	status := service.AddClanService(interaction, disciplineIDsSuffix)["status"]
	ResolveResponse(s, interaction, status)
}

package controller

import (
	"vtm-go-bot/model"
	"vtm-go-bot/service"

	"github.com/bwmarrin/discordgo"
)

func AddMerit(s *discordgo.Session, interaction *discordgo.InteractionCreate) string {
	status := service.AddMeritService(interaction, "")["status"]
	return status
}

func GetMeritByID(meritID string) model.Merit {
	return service.GetMeritByID(meritID)
}

func GetMeritsByKind(meritKind string) []model.Merit {
	return service.GetMeritsByKind(meritKind)
}

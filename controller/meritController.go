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

func GetMeritKindName(kindID string) string {
	return service.GetMeritKindName(kindID)
}

func UpdateMerit(s *discordgo.Session, interaction *discordgo.InteractionCreate, meritID string) string {
	status := service.AddMeritService(interaction, meritID)["status"]
	return status
}

func DeleteMerit(s *discordgo.Session, interaction *discordgo.InteractionCreate, meritID string) string {
	status := service.DeleteMeritService(interaction, meritID)["status"]
	return status
}

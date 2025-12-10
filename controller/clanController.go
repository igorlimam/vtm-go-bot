package controller

import (
	"vtm-go-bot/model"
	"vtm-go-bot/service"

	"github.com/bwmarrin/discordgo"
)

func AddClan(s *discordgo.Session, interaction *discordgo.InteractionCreate, disciplineIDsSuffix string) string {
	status := service.AddClanService(interaction, disciplineIDsSuffix)["status"]
	return status
}

func GetAllClans() []model.Clan {
	return service.GetAllClansService()
}

func GetClanByID(idStr string) model.Clan {
	return service.GetClanByIDService(idStr)
}

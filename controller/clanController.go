package controller

import (
	"vtm-go-bot/model"
	"vtm-go-bot/service"

	"github.com/bwmarrin/discordgo"
)

func AddClan(s *discordgo.Session, interaction *discordgo.InteractionCreate, disciplineIDsSuffix string) string {
	isUpdate := false
	status := service.AddClanService(interaction, disciplineIDsSuffix, isUpdate)["status"]
	return status
}

func GetAllClans() []model.Clan {
	return service.GetAllClansService()
}

func GetClanByID(idStr string) model.Clan {
	return service.GetClanByIDService(idStr)
}

func UpdateClan(s *discordgo.Session, interaction *discordgo.InteractionCreate, disciplineIDsSuffix string) string {
	isUpdate := true
	status := service.AddClanService(interaction, disciplineIDsSuffix, isUpdate)["status"]
	return status
}

func DeleteClan(s *discordgo.Session, interaction *discordgo.InteractionCreate, clanID string) string {
	status := service.DeleteClanService(interaction, clanID)["status"]
	return status
}

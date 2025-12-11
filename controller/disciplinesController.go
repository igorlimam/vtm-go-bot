package controller

import (
	"log"
	"vtm-go-bot/model"
	"vtm-go-bot/service"

	"github.com/bwmarrin/discordgo"
)

func AddDiscipline(session *discordgo.Session, interaction *discordgo.InteractionCreate) string {
	log.Println("AddDiscipline command invoked")
	status := service.AddDisciplineService(interaction, "")["status"]
	return status
}

func UpdateDiscipline(session *discordgo.Session, interaction *discordgo.InteractionCreate, id string) string {
	log.Println("UpdateDiscipline command invoked")
	status := service.AddDisciplineService(interaction, id)["status"]
	return status
}

func GetAllDisciplines() []model.Discipline {
	return service.GetAllDisciplines()
}

func GetDisciplineByID(id string) model.Discipline {
	discipline := service.GetDisciplineByID(id)
	return discipline
}

func DeleteDiscipline(s *discordgo.Session, interaction *discordgo.InteractionCreate, idStr string) string {
	log.Println("DeleteDiscipline command invoked")
	status := service.DeleteDiscipline(idStr)
	return status
}

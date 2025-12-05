package service

import (
	"log"
	"vtm-go-bot/model"
	"vtm-go-bot/repository"

	"github.com/bwmarrin/discordgo"
)

func AddDisciplineService(interaction *discordgo.InteractionCreate) map[string]string {
	dataModal := ModalToMap(interaction)
	status := repository.AddDiscipline(
		dataModal["discipline-name"].(string),
		dataModal["discipline-type"].(string),
		dataModal["discipline-resonance"].(string),
		dataModal["discipline-threat"].(string),
		dataModal["discipline-description"].(string),
	)

	log.Printf("Inserted Discipline: %s\n", dataModal["discipline-name"].(string))
	return status
}

func GetAllDisciplines() []model.Discipline {
	disciplines := repository.GetAllDisciplines()
	return disciplines
}

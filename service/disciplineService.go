package service

import (
	"log"
	"vtm-go-bot/repository"

	"github.com/bwmarrin/discordgo"
)

func AddDisciplineService(interaction *discordgo.InteractionCreate) map[string]string {
	//log.Printf("Interaction Data: %+v\n", interaction.Data)
	dataModal := ModalToMap(interaction)

	//log.Printf("Modal Data: %+v\n", dataModal)
	//log.Printf("%s, %s, %s, %s", dataModal["discipline-name"], dataModal["discipline-type"], dataModal["discipline-resonance"], dataModal["discipline-threat"])

	repository.AddDiscipline(
		dataModal["discipline-name"].(string),
		dataModal["discipline-type"].(string),
		dataModal["discipline-resonance"].(string),
		dataModal["discipline-threat"].(string),
		dataModal["discipline-description"].(string),
	)

	log.Printf("Inserted Discipline: %s\n", dataModal["discipline-name"].(string))
	return map[string]string{"status": "Disciplina adicionada com sucesso!"}
}

package service

import (
	"log"
	"vtm-go-bot/repository"

	"github.com/bwmarrin/discordgo"
)

func AddDisciplineService(interaction *discordgo.InteractionCreate) map[string]string {
	//log.Printf("Interaction Data: %+v\n", interaction.Data)
	dataModal := ModalToMap(interaction)

	log.Printf("Modal Data: %+v\n", dataModal)
	//data := interaction.ApplicationCommandData()
	log.Printf("%s, %s, %s, %s", dataModal["name"], dataModal["type"], dataModal["resonance"], dataModal["threat"])
	discipline := repository.AddDiscipline(dataModal["name"], dataModal["type"], dataModal["resonance"], dataModal["threat"])
	log.Printf("Inserted Discipline: %+v\n", discipline)
	return map[string]string{"status": "Disciplina adicionada com sucesso!"}
}

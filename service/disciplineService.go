package service

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func AddDisciplineService(interaction *discordgo.InteractionCreate) map[string]string {
	//log.Printf("Interaction Data: %+v\n", interaction.Data)
	dataModal := ModalToMap(interaction)

	log.Printf("Modal Data: %+v\n", dataModal)
	//data := interaction.ApplicationCommandData()

	// Business logic to add a discipline would go here
	return dataModal
}

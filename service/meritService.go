package service

import (
	"vtm-go-bot/repository"

	"github.com/bwmarrin/discordgo"
)

func AddMeritService(interaction *discordgo.InteractionCreate, meritID string) map[string]string {
	dataModal := ModalToMap(interaction)

	var status map[string]string
	if meritID == "" {
		status = repository.AddMerit(
			dataModal["merit-name"].(string),
			dataModal["merit-description"].(string),
			dataModal["merit-kind"].(string),
			dataModal["merit-levels-info"].(string),
		)
	} else {
		status = map[string]string{"status": "Atualização de mérito não implementada."}
	}

	return status
}

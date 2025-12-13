package service

import (
	"log"
	"strconv"
	"vtm-go-bot/model"
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

func GetMeritByID(meritID string) model.Merit {
	id, err := strconv.Atoi(meritID)
	if err != nil {
		log.Printf("Invalid merit ID: %v", err)
		return model.Merit{}
	}
	return repository.GetMeritByID(uint(id))
}

func GetMeritsByKind(meritKind string) []model.Merit {
	return repository.GetMeritsByKind(meritKind)
}

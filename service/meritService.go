package service

import (
	"log"
	"slices"
	"strconv"
	"vtm-go-bot/model"
	"vtm-go-bot/repository"

	"github.com/bwmarrin/discordgo"
)

func GetMeritKindName(kindID string) string {
	switch kindID {
	case "1":
		return "Vantagem"
	case "2":
		return "Desvantagem"
	case "3":
		return "Antecedente"
	default:
		return "Desconhecido"
	}
}

func AddMeritService(interaction *discordgo.InteractionCreate, meritID string) map[string]string {
	dataModal := ModalToMap(interaction)

	meritKind := dataModal["merit-kind"].(string)
	kinds := []string{"Vantagem", "Desvantagem", "Antecedente"}
	if !slices.Contains(kinds, meritKind) {
		return map[string]string{"status": "MÉRITO NÃO CADASTRADO! Tipo de mérito inválido!"}
	}

	var status map[string]string
	if meritID == "" {
		status = repository.AddMerit(
			dataModal["merit-name"].(string),
			dataModal["merit-description"].(string),
			meritKind,
			dataModal["merit-levels-info"].(string),
		)
	} else {
		id, err := strconv.Atoi(meritID)
		if err != nil {
			log.Printf("Invalid merit ID: %v", err)
			return map[string]string{"status": "MÉRITO NÃO ATUALIZADO! ID inválido."}
		}
		status = repository.UpdateMerit(
			uint(id),
			dataModal["merit-name"].(string),
			dataModal["merit-description"].(string),
			meritKind,
			dataModal["merit-levels-info"].(string),
		)
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

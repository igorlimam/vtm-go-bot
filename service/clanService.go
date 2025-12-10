package service

import (
	"log"
	"strconv"
	"strings"
	"vtm-go-bot/model"
	"vtm-go-bot/repository"

	"github.com/bwmarrin/discordgo"
)

func AddClanService(interaction *discordgo.InteractionCreate, disciplineIDsSuffix string) map[string]string {

	dataModal := ModalToMap(interaction)
	disciplinesSplited := strings.Split(disciplineIDsSuffix, "-")
	var disciplinesVector []model.Discipline

	for _, idStr := range disciplinesSplited {
		id, _ := strconv.Atoi(idStr)
		disciplinesVector = append(disciplinesVector, repository.GetDisciplineById(uint(id)))
	}

	repository.AddClan(
		dataModal["clan-name"].(string),
		dataModal["clan-description"].(string),
		dataModal["clan-bane"].(string),
		dataModal["clan-compulsion"].(string),
		disciplinesVector,
	)

	return map[string]string{"status": "Clan added successfully"}
}

func GetAllClansService() []model.Clan {
	clans := repository.GetAllClans()
	return clans
}

func GetClanByIDService(idStr string) model.Clan {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Failed to parse id '%s': %v", idStr, err)
		return model.Clan{}
	}
	clan := repository.GetClanByID(uint(id))
	return clan
}

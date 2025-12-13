package service

import (
	"log"
	"strconv"
	"strings"
	"vtm-go-bot/model"
	"vtm-go-bot/repository"

	"github.com/bwmarrin/discordgo"
)

func AddClanService(interaction *discordgo.InteractionCreate, disciplineIDsSuffix string, isUpdate bool) map[string]string {
	dataModal := ModalToMap(interaction)
	disciplinesSplited := strings.Split(disciplineIDsSuffix, "-")
	var clanId uint64
	var disciplinesVector []model.Discipline

	for i, idStr := range disciplinesSplited {
		id, _ := strconv.Atoi(idStr)
		if i == 0 && isUpdate {
			clanId = uint64(id)
		} else {
			disciplinesVector = append(disciplinesVector, repository.GetDisciplineById(uint(id)))
		}
	}

	var status map[string]string
	if isUpdate {
		status = repository.UpdateClan(
			uint(clanId),
			dataModal["clan-name"].(string),
			dataModal["clan-description"].(string),
			dataModal["clan-bane"].(string),
			dataModal["clan-compulsion"].(string),
			disciplinesVector,
		)

		log.Printf("Updated Clan ID %d: %s\n", clanId, dataModal["clan-name"].(string))
	} else {
		status = repository.AddClan(
			dataModal["clan-name"].(string),
			dataModal["clan-description"].(string),
			dataModal["clan-bane"].(string),
			dataModal["clan-compulsion"].(string),
			disciplinesVector,
		)
	}

	return status
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

func DeleteClanService(interaction *discordgo.InteractionCreate, clanID string) map[string]string {
	id, err := strconv.Atoi(clanID)
	if err != nil {
		log.Printf("Failed to parse clanID '%s': %v", clanID, err)
		return map[string]string{"status": "Erro ao deletar clã: ID inválido."}
	}
	status := repository.DeleteClan(uint(id))
	log.Printf("Deleted Clan ID %d\n", id)
	return status
}

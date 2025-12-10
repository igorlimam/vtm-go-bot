package service

import (
	"log"
	"strconv"
	"vtm-go-bot/model"
	"vtm-go-bot/repository"

	"github.com/bwmarrin/discordgo"
)

func AddDisciplineService(interaction *discordgo.InteractionCreate, idStr string) map[string]string {
	dataModal := ModalToMap(interaction)

	var status map[string]string

	if idStr != "" {
		// If idStr is provided, we are updating an existing discipline
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return map[string]string{"status": "DISCIPLINA NÃO ATUALIZADA! ID inválido."}
		}
		status = repository.UpdateDiscipline(
			uint(id),
			dataModal["discipline-name"].(string),
			dataModal["discipline-type"].(string),
			dataModal["discipline-resonance"].(string),
			dataModal["discipline-threat"].(string),
			dataModal["discipline-description"].(string),
		)

		log.Printf("Updated Discipline ID %d: %s\n", id, dataModal["discipline-name"].(string))
	} else {
		status = repository.AddDiscipline(
			dataModal["discipline-name"].(string),
			dataModal["discipline-type"].(string),
			dataModal["discipline-resonance"].(string),
			dataModal["discipline-threat"].(string),
			dataModal["discipline-description"].(string),
		)

		log.Printf("Inserted Discipline: %s\n", dataModal["discipline-name"].(string))
	}

	return status
}

func GetDisciplineByID(idStr string) model.Discipline {
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Printf("Failed to parse id '%s': %v", idStr, err)
		return model.Discipline{}
	}
	discipline := repository.GetDisciplineById(uint(id))
	return discipline
}

func GetAllDisciplines() []model.Discipline {
	disciplines := repository.GetAllDisciplines()
	return disciplines
}

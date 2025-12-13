package service

import (
	"strconv"
	"vtm-go-bot/model"
	"vtm-go-bot/repository"

	"github.com/bwmarrin/discordgo"
)

func AddPowerService(interaction *discordgo.InteractionCreate, disciplineId string, powerId string) map[string]string {
	dataModal := ModalToMap(interaction)
	disciplineIdInt, _ := strconv.Atoi(disciplineId)

	nameLevelType := SplitModalInput(dataModal["power-name-level-type"].(string), "|", 3)
	name := nameLevelType[0]
	levelStr := nameLevelType[1]
	powerType := nameLevelType[2]
	costDuration := SplitModalInput(dataModal["power-cost-duration-amalgam"].(string), "|", 3)
	cost := costDuration[0]
	duration := costDuration[1]
	amalgam := costDuration[2]

	level := ConvertStringToInt(levelStr)
	if level < 1 || level > 10 {
		return map[string]string{"status": "PODER NÃO CADASTRADO! Level inválido! Deve ser entre 1 e 10."}
	}

	var status map[string]string

	if powerId != "" {
		// Updating existing power
		powerIdInt, err := strconv.Atoi(powerId)
		if err != nil {
			return map[string]string{"status": "PODER NÃO ATUALIZADO! ID inválido."}
		}
		status = repository.UpdatePower(
			uint(powerIdInt),
			uint(disciplineIdInt),
			name,
			dataModal["power-description"].(string),
			dataModal["power-dice-pool"].(string),
			cost,
			duration,
			dataModal["power-system"].(string),
			powerType,
			amalgam,
			level,
		)
	} else {
		status = repository.AddPower(
			uint(disciplineIdInt),
			name,
			dataModal["power-description"].(string),
			dataModal["power-dice-pool"].(string),
			cost,
			duration,
			dataModal["power-system"].(string),
			powerType,
			amalgam,
			level,
		)
	}

	return status
}

func GetAllPowers() []model.Power {
	return repository.GetAllPowers()
}

func GetDisciplinePowersByID(disciplineID string) []model.Power {
	id, err := strconv.Atoi(disciplineID)
	if err != nil {
		return []model.Power{}
	}
	power := repository.GetPowersByDiciplineId(uint(id))
	return power
}

func GetPowerById(powerID string) model.Power {
	id, err := strconv.Atoi(powerID)
	if err != nil {
		return model.Power{}
	}
	power := repository.GetPowerById(uint(id))
	return power
}

func DeletePowerService(interaction *discordgo.InteractionCreate, powerID string) map[string]string {
	id, err := strconv.Atoi(powerID)
	if err != nil {
		return map[string]string{"status": "Erro ao deletar poder: ID inválido."}
	}
	status := repository.DeletePower(uint(id))
	return status
}

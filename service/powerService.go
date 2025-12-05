package service

import (
	"strconv"
	"vtm-go-bot/repository"

	"github.com/bwmarrin/discordgo"
)

func AddPowerService(interaction *discordgo.InteractionCreate, disciplineId string) map[string]string {
	dataModal := ModalToMap(interaction)
	disciplineIdUint, _ := strconv.ParseUint(disciplineId, 10, 64)
	level := ConvertStringToInt(dataModal["power-level"].(string))
	if level < 1 || level > 10 {
		return map[string]string{"status": "PODER NÃO CADASTRADO! Level inválido! Deve ser entre 1 e 10."}
	}

	status := repository.AddPower(
		uint(disciplineIdUint),
		dataModal["power-name"].(string),
		dataModal["power-description"].(string),
		dataModal["power-dice-pool"].(string),
		dataModal["power-cost"].(string),
		dataModal["power-duration"].(string),
		dataModal["power-system"].(string),
		dataModal["power-type"].(string),
		level,
	)

	return status
}

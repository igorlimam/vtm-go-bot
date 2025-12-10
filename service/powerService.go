package service

import (
	"strconv"
	"vtm-go-bot/repository"

	"github.com/bwmarrin/discordgo"
)

func AddPowerService(interaction *discordgo.InteractionCreate, disciplineId string) map[string]string {
	dataModal := ModalToMap(interaction)
	disciplineIdUint, _ := strconv.ParseUint(disciplineId, 10, 64)

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

	status := repository.AddPower(
		uint(disciplineIdUint),
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

	return status
}

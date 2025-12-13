package view

import (
	"fmt"
	"strconv"
	"strings"
	"vtm-go-bot/model"

	"github.com/bwmarrin/discordgo"
)

func PowerSelectDisciplineView(s *discordgo.Session, interaction *discordgo.InteractionCreate, disciplines []model.Discipline) {
	customID := "select-discipline-for-power"
	placeholder := "Selecione a disciplina para o novo poder"
	contentPlaceholder := "Escolha a disciplina para o novo poder:"

	var options []map[string]string

	for _, discipline := range disciplines {
		options = append(options, map[string]string{
			"label": discipline.Name,
			"value": fmt.Sprintf("%d", discipline.ID),
		})
	}

	SelectMenu(s, interaction, options, customID, placeholder, contentPlaceholder, 1)
}

func AddPowerView(s *discordgo.Session, interaction *discordgo.InteractionCreate, disciplineId string, power *model.Power) {
	customID := "add-power-modal|" + disciplineId
	title := "Adicionar Novo Poder"
	nameLevelKindValue := ""
	descriptionValue := ""
	dicePoolValue := ""
	costDurationAmalgamValue := ""
	systemValue := ""

	if power != nil {
		customID = "update-power-modal|" + fmt.Sprintf("%d", power.ID) + "|" + disciplineId
		title = "Atualizar Poder"
		nameLevelKindValue = fmt.Sprintf("%s|%d|%s", power.Name, power.Level, power.Kind)
		descriptionValue = power.Description
		dicePoolValue = power.DicePool
		costDurationAmalgamValue = fmt.Sprintf("%s|%s|%s", power.Cost, power.Duration, power.Amalgam)
		systemValue = power.System
	}

	fields := []map[string]string{
		{
			"customID": "power-name-level-type",
			"label":    "Nome|Nivel|Tipo - \"Fata Morgana|1|Mental\"",
			"style":    "short",
			"value":    nameLevelKindValue,
		},
		{
			"customID": "power-description",
			"label":    "Descrição do Poder",
			"style":    "paragraph",
			"value":    descriptionValue,
		},
		{
			"customID": "power-dice-pool",
			"label":    "Parada de Dados",
			"style":    "short",
			"value":    dicePoolValue,
		},
		{
			"customID": "power-cost-duration-amalgam",
			"label":    "Custo|Duração|Amalgama",
			"style":    "short",
			"value":    costDurationAmalgamValue,
		},
		{
			"customID": "power-system",
			"label":    "Sistema",
			"style":    "paragraph",
			"value":    systemValue,
		},
	}

	Modal(s, interaction, customID, title, fields)
}

func PowerInfoView(s *discordgo.Session, interaction *discordgo.InteractionCreate, query string, powers []model.Power) {

	options := []map[string]string{}
	for _, power := range powers {
		if strings.Contains(strings.ToLower(power.Name), query) && (len(options) < 25) {
			options = append(options, map[string]string{
				"label": power.Name,
				"value": fmt.Sprintf("%d", power.ID),
			})
		}
	}

	AutoComplete(s, interaction, options)
}

func ShowPowerInfoView(s *discordgo.Session, interaction *discordgo.InteractionCreate, power model.Power) {

	embedFields := []map[string]string{}

	embedFields = append(embedFields, map[string]string{
		"name":   "Descrição",
		"value":  power.Description,
		"inline": "false",
	})
	embedFields = append(embedFields, map[string]string{
		"name":   "Parada de dados",
		"value":  power.DicePool,
		"inline": "true",
	})
	embedFields = append(embedFields, map[string]string{
		"name":   "Custo",
		"value":  power.Cost,
		"inline": "true",
	})
	embedFields = append(embedFields, map[string]string{
		"name":   "Duração",
		"value":  power.Duration,
		"inline": "false",
	})
	embedFields = append(embedFields, map[string]string{
		"name":   "Tipo",
		"value":  power.Kind,
		"inline": "true",
	})
	embedFields = append(embedFields, map[string]string{
		"name":   "Amalgama",
		"value":  power.Amalgam,
		"inline": "true",
	})
	embedFields = append(embedFields, map[string]string{
		"name":   "Nivel",
		"value":  strconv.Itoa(power.Level),
		"inline": "true",
	})
	embedFields = append(embedFields, map[string]string{
		"name":   "Sistema",
		"value":  power.System,
		"inline": "false",
	})

	EmbedMessage(s, interaction, embedFields, power.Name, "")
}

func ConfirmDeletePower(s *discordgo.Session, interaction *discordgo.InteractionCreate, power model.Power, disciplineName string) {

	customIDConfirmation := "confirm-delete-power|" + fmt.Sprintf("%d", power.ID)
	customIDCancel := "cancel-delete-power"
	messageContent := fmt.Sprintf("Tem certeza que deseja deletar o poder **%s** da disciplina **%s**?", power.Name, disciplineName)

	ConfirmationButton(s, interaction, customIDConfirmation, customIDCancel, messageContent)
}

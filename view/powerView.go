package view

import (
	"fmt"
	"log"
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

	choices := []*discordgo.ApplicationCommandOptionChoice{}
	for _, power := range powers {
		if strings.Contains(strings.ToLower(power.Name), query) && (len(choices) < 25) {
			choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
				Name:  power.Name,
				Value: fmt.Sprintf("%d", power.ID),
			})
		}
	}
	err := s.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionApplicationCommandAutocompleteResult,
		Data: &discordgo.InteractionResponseData{
			Choices: choices,
		},
	})

	if err != nil {
		log.Println("Error responding to autocomplete POWER interaction:", err)
	}
}

func ShowPowerInfoView(s *discordgo.Session, interaction *discordgo.InteractionCreate, power model.Power) {
	embed := &discordgo.MessageEmbed{
		Title:       power.Name,
		Description: power.Description,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Parada de dados",
				Value:  power.DicePool,
				Inline: true,
			},
			{
				Name:   "Custo",
				Value:  power.Cost,
				Inline: true,
			},
			{
				Name:   "Duração",
				Value:  power.Duration,
				Inline: false,
			},
			{
				Name:   "Tipo",
				Value:  power.Kind,
				Inline: true,
			},
			{
				Name:   "Amalgama",
				Value:  power.Amalgam,
				Inline: true,
			},
			{
				Name:   "Nivel",
				Value:  strconv.Itoa(power.Level),
				Inline: true,
			},
			{
				Name:   "Sistema",
				Value:  power.System,
				Inline: false,
			},
		},
	}
	s.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})
}

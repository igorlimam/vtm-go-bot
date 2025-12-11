package view

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"vtm-go-bot/model"

	"github.com/bwmarrin/discordgo"
)

func PowerSelectDisciplineView(s *discordgo.Session, interaction *discordgo.InteractionCreate, disciplines []model.Discipline) error {
	choices := make([]discordgo.SelectMenuOption, 0)
	for _, discipline := range disciplines {
		choices = append(choices, discordgo.SelectMenuOption{
			Label: discipline.Name,
			Value: fmt.Sprintf("%d", discipline.ID),
		})
	}

	return s.InteractionRespond(
		interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Escolha a disciplina para o novo poder:",
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.SelectMenu{
								CustomID: "select-discipline-for-power",
								Options:  choices,
							},
						},
					},
				},
			},
		},
	)

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
	err := s.InteractionRespond(
		interaction.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseModal,
			Data: &discordgo.InteractionResponseData{
				CustomID: customID,
				Title:    title,
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID: "power-name-level-type",
								Label:    "Nome|Nivel|Tipo - \"Fata Morgana|1|Mental\"",
								Style:    discordgo.TextInputShort,
								Value:    nameLevelKindValue,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID: "power-description",
								Label:    "Descrição do Poder",
								Style:    discordgo.TextInputParagraph,
								Value:    descriptionValue,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID: "power-dice-pool",
								Label:    "Parada de Dados",
								Style:    discordgo.TextInputShort,
								Value:    dicePoolValue,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{

							discordgo.TextInput{
								CustomID: "power-cost-duration-amalgam",
								Label:    "Custo|Duração|Amalgama",
								Style:    discordgo.TextInputShort,
								Value:    costDurationAmalgamValue,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID: "power-system",
								Label:    "Sistema",
								Style:    discordgo.TextInputParagraph,
								Value:    systemValue,
							},
						},
					},
				},
			},
		},
	)
	if err != nil {
		fmt.Println("Error responding with modal:", err)
	}
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

package view

import (
	"fmt"
	"log"
	"slices"
	"strings"
	"vtm-go-bot/model"

	"github.com/bwmarrin/discordgo"
)

func StringSelectClanDisciplines(s *discordgo.Session, interaction *discordgo.InteractionCreate, disciplines []model.Discipline, selectedDisciplines []model.Discipline, clanID string) {

	var options []map[string]string

	for _, discipline := range disciplines {
		var def string
		if selectedDisciplines == nil {
			def = "false"
		} else {
			def = fmt.Sprintf("%t", slices.ContainsFunc(selectedDisciplines, func(d model.Discipline) bool { return d.ID == discipline.ID }))
		}
		options = append(options, map[string]string{
			"label":   discipline.Name,
			"value":   fmt.Sprintf("%d", discipline.ID),
			"default": def,
		})
	}

	if clanID == "" {
		clanID = "0"
	}
	customID := "select-disciplines-for-clan|" + clanID
	placeholder := "Selecione pelo menos uma disciplina"
	contentPlaceholder := "Escolha as disciplinas para o novo clã:"

	SelectMenu(s, interaction, options, customID, placeholder, contentPlaceholder)
}

func AddClanView(s *discordgo.Session, interaction *discordgo.InteractionCreate, disciplineIDs []string, clan *model.Clan) {
	suffix := strings.Join(disciplineIDs, "-")

	customID := "add-clan-modal|" + suffix
	title := "Adicionar Novo Clã"
	name := ""
	description := ""
	bane := ""
	compulsion := ""

	if clan != nil {
		customID = "update-clan-modal|" + fmt.Sprintf("%d", clan.ID) + "-" + suffix
		title = "Atualizar Clã"
		name = clan.Name
		description = clan.Description
		bane = clan.Bane
		compulsion = clan.Compulsion
	}

	fields := []map[string]string{
		{
			"customID": "clan-name",
			"label":    "Nome do Clã",
			"value":    name,
			"style":    "short",
		},
		{
			"customID": "clan-description",
			"label":    "Descrição do Clã",
			"value":    description,
			"style":    "paragraph",
		},
		{
			"customID": "clan-bane",
			"label":    "Fraqueza do Clã",
			"value":    bane,
			"style":    "paragraph",
		},
		{
			"customID": "clan-compulsion",
			"label":    "Compulsão do Clã",
			"value":    compulsion,
			"style":    "paragraph",
		},
	}

	Modal(s, interaction, customID, title, fields)
}

func ClanInfoView(s *discordgo.Session, interaction *discordgo.InteractionCreate, clans []model.Clan) {
	var choices []*discordgo.ApplicationCommandOptionChoice
	for _, clan := range clans {
		choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
			Name:  clan.Name,
			Value: fmt.Sprintf("%d", clan.ID),
		})
	}

	err := s.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionApplicationCommandAutocompleteResult,
		Data: &discordgo.InteractionResponseData{
			Choices: choices,
		},
	})

	if err != nil {
		log.Println("Error responding to autocomplete interaction:", err)
	}
}

func ShowClanInfoView(s *discordgo.Session, interaction *discordgo.InteractionCreate, clan model.Clan, disciplines []model.Discipline) {

	disciplinesList := ""
	for _, discipline := range disciplines {
		disciplinesList += fmt.Sprintf("- %s\n", discipline.Name)
	}

	embed := &discordgo.MessageEmbed{
		Title:       clan.Name,
		Description: clan.Description,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Fraqueza",
				Value:  clan.Bane,
				Inline: false,
			},
			{
				Name:   "Compulsão",
				Value:  clan.Compulsion,
				Inline: false,
			},
			{
				Name:   "Disciplinas",
				Value:  disciplinesList,
				Inline: true,
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

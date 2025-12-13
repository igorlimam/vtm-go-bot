package view

import (
	"fmt"
	"slices"
	"strings"
	"vtm-go-bot/model"

	"github.com/bwmarrin/discordgo"
)

func StringSelectClanDisciplines(s *discordgo.Session, interaction *discordgo.InteractionCreate, disciplines []model.Discipline, selectedDisciplines []model.Discipline, clanID string) {
	maxDisciplines := 3
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

	SelectMenu(s, interaction, options, customID, placeholder, contentPlaceholder, maxDisciplines)
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

	options := []map[string]string{}
	for _, clan := range clans {
		options = append(options, map[string]string{
			"label": clan.Name,
			"value": fmt.Sprintf("%d", clan.ID),
		})
	}

	AutoComplete(s, interaction, options)
}

func ShowClanInfoView(s *discordgo.Session, interaction *discordgo.InteractionCreate, clan model.Clan, disciplines []model.Discipline) {

	disciplinesList := ""
	for _, discipline := range disciplines {
		disciplinesList += fmt.Sprintf("- %s\n", discipline.Name)
	}

	embedFields := []map[string]string{}

	embedFields = append(embedFields, map[string]string{
		"name":   "Fraqueza",
		"value":  clan.Bane,
		"inline": "false",
	})
	embedFields = append(embedFields, map[string]string{
		"name":   "Compulsão",
		"value":  clan.Compulsion,
		"inline": "false",
	})
	embedFields = append(embedFields, map[string]string{
		"name":   "Disciplinas",
		"value":  disciplinesList,
		"inline": "false",
	})
	EmbedMessage(s, interaction, embedFields, clan.Name, clan.Description)
}

func ConfirmDeleteClan(s *discordgo.Session, interaction *discordgo.InteractionCreate, clan model.Clan) {

	customIDConfirmation := "confirm-delete-clan|" + fmt.Sprintf("%d", clan.ID)
	customIDCancel := "cancel-delete-clan"
	messageContent := fmt.Sprintf("Tem certeza que deseja deletar o clã **%s**?", clan.Name)

	ConfirmationButton(s, interaction, customIDConfirmation, customIDCancel, messageContent)
}

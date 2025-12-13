package view

import (
	"fmt"
	"strings"
	"vtm-go-bot/model"

	"github.com/bwmarrin/discordgo"
)

func AddMeritView(s *discordgo.Session, interaction *discordgo.InteractionCreate, meritKind string, merit *model.Merit) {
	customID := "add-merit-modal|0"
	title := "Adicionar Mérito"

	name := ""
	description := ""
	kind := meritKind
	levelsInfo := ""

	if merit != nil {
		customID = "update-merit-modal|" + fmt.Sprintf("%d", merit.ID)
		title = "Atualizar Mérito"
		name = merit.Name
		description = merit.Description
		kind = merit.Kind
		levelsInfo = merit.LevelsInfo
	}

	fields := []map[string]string{
		{
			"customID": "merit-name",
			"label":    "Nome",
			"value":    name,
			"style":    "short",
		},
		{
			"customID": "merit-description",
			"label":    "Descrição",
			"value":    description,
			"style":    "paragraph",
		},
		{
			"customID": "merit-kind",
			"label":    "Tipo - Vantagem, Desvantagem, Antecedente",
			"value":    kind,
			"style":    "short",
		},
		{
			"customID": "merit-levels-info",
			"label":    "Informações de Níveis (se aplicável)",
			"value":    levelsInfo,
			"style":    "paragraph",
		},
	}

	Modal(s, interaction, customID, title, fields)
}

func StringSelectMeritKindView(s *discordgo.Session, interaction *discordgo.InteractionCreate) {
	customID := "select-merit-kind"
	placeholder := "Selecione o tipo de mérito"
	contentPlaceholder := "Escolha o tipo de mérito que deseja adicionar:"

	options := []map[string]string{
		{
			"label": "Vantagem",
			"value": "1",
		},
		{
			"label": "Desvantagem",
			"value": "2",
		},
		{
			"label": "Antecedente",
			"value": "3",
		},
	}

	SelectMenu(s, interaction, options, customID, placeholder, contentPlaceholder, 1)
}

func MeritKindInfoView(s *discordgo.Session, interaction *discordgo.InteractionCreate) {

	meritKinds := []map[string]string{
		{"label": "Vantagem", "value": "1"},
		{"label": "Desvantagem", "value": "2"},
		{"label": "Antecedente", "value": "3"},
	}

	options := []map[string]string{}
	for _, meritKind := range meritKinds {
		options = append(options, map[string]string{
			"label": meritKind["label"],
			"value": meritKind["value"],
		})
	}

	AutoComplete(s, interaction, options)
}

func MeritInfoView(s *discordgo.Session, interaction *discordgo.InteractionCreate, query string, merits []model.Merit) {
	options := []map[string]string{}
	for _, merit := range merits {
		cond := query == "" || strings.Contains(strings.ToLower(merit.Name), query)
		if cond && (len(options) < 25) {
			options = append(options, map[string]string{
				"label": merit.Name,
				"value": fmt.Sprintf("%d", merit.ID),
			})
		}
	}

	AutoComplete(s, interaction, options)
}

func ShowMeritInfoView(s *discordgo.Session, interaction *discordgo.InteractionCreate, merit model.Merit) {

	title := merit.Name
	description := merit.Description

	embedFields := []map[string]string{}

	embedFields = append(embedFields, map[string]string{
		"name":   "Tipo",
		"value":  merit.Kind,
		"inline": "true",
	})
	if merit.LevelsInfo != "" {
		embedFields = append(embedFields, map[string]string{
			"name":   "Informações de Níveis",
			"value":  merit.LevelsInfo,
			"inline": "false",
		})
	}

	EmbedMessage(s, interaction, embedFields, title, description)
}

func ConfirmDeleteMerit(s *discordgo.Session, interaction *discordgo.InteractionCreate, meritKind string, merit model.Merit) {

	if meritKind != "Antecedente" {
		meritKind = "a " + meritKind
	} else {
		meritKind = "o " + meritKind
	}

	content := fmt.Sprintf("Você tem certeza que deseja deletar %s **%s**?", meritKind, merit.Name)
	confirmCustomID := fmt.Sprintf("confirm-delete-merit|%d", merit.ID)
	cancelCustomID := "cancel-delete-merit"

	ConfirmationButton(s, interaction, confirmCustomID, cancelCustomID, content)
}

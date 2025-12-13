package view

import (
	"fmt"
	"vtm-go-bot/model"

	"github.com/bwmarrin/discordgo"
)

func AddMeritView(s *discordgo.Session, interaction *discordgo.InteractionCreate, meritKind string, merit *model.Merit) {
	customID := "add-merit-modal|0"
	title := "Adicionar Mérito"

	switch meritKind {
	case "1":
		meritKind = "Vantagem"
	case "2":
		meritKind = "Desvantagem"
	case "3":
		meritKind = "Antecedente"
	}

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

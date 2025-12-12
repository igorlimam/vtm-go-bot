package view

import (
	"fmt"
	"vtm-go-bot/model"

	"github.com/bwmarrin/discordgo"
)

func AddDisciplineView(s *discordgo.Session, interaction *discordgo.InteractionCreate, discipline *model.Discipline) {
	customID := "add-discipline-modal|0"
	title := "Adicionar Nova Disciplina"

	// prefill values if updating
	nameValue := ""
	typeValue := ""
	resonanceValue := ""
	descriptionValue := ""
	threatValue := ""

	if discipline != nil {
		customID = "update-discipline-modal|" + fmt.Sprintf("%d", discipline.ID)
		title = "Atualizar Disciplina"
		nameValue = discipline.Name
		typeValue = discipline.Dtype
		resonanceValue = discipline.Resonance
		descriptionValue = discipline.Description
		threatValue = discipline.Threat
	}

	fields := []map[string]string{
		{
			"customID": "discipline-name",
			"label":    "Nome da Disciplina",
			"style":    "short",
			"value":    nameValue,
		},
		{
			"customID": "discipline-type",
			"label":    "Tipo da Disciplina",
			"style":    "short",
			"value":    typeValue,
		},
		{
			"customID": "discipline-resonance",
			"label":    "Ressonância da Disciplina",
			"style":    "short",
			"value":    resonanceValue,
		},
		{
			"customID": "discipline-description",
			"label":    "Descrição da Disciplina",
			"style":    "paragraph",
			"value":    descriptionValue,
		},
		{
			"customID": "discipline-threat",
			"label":    "Ameaça da Disciplina",
			"style":    "paragraph",
			"value":    threatValue,
		},
	}

	Modal(s, interaction, customID, title, fields)
}

func DisciplinaInfoView(s *discordgo.Session, interaction *discordgo.InteractionCreate, disciplines []model.Discipline) {

	options := []map[string]string{}
	for _, discipline := range disciplines {
		options = append(options, map[string]string{
			"label": discipline.Name,
			"value": fmt.Sprintf("%d", discipline.ID),
		})
	}

	AutoComplete(s, interaction, options)
}

func ShowDisciplineInfoView(s *discordgo.Session, interaction *discordgo.InteractionCreate, discipline model.Discipline) {

	embedFields := []map[string]string{}

	embedFields = append(embedFields, map[string]string{
		"name":   "Tipo",
		"value":  discipline.Dtype,
		"inline": "true",
	})
	embedFields = append(embedFields, map[string]string{
		"name":   "Ressonância",
		"value":  discipline.Resonance,
		"inline": "true",
	})
	embedFields = append(embedFields, map[string]string{
		"name":   "Ameaça",
		"value":  discipline.Threat,
		"inline": "false",
	})

	EmbedMessage(s, interaction, embedFields, discipline.Name, discipline.Description)
}

func ConfirmDeleteDiscipline(s *discordgo.Session, interaction *discordgo.InteractionCreate, discipline model.Discipline) {

	customIDConfirmation := "confirm-delete-discipline|" + fmt.Sprintf("%d", discipline.ID)
	customIDCancel := "cancel-delete-discipline"
	messageContent := fmt.Sprintf("Tem certeza que deseja deletar a disciplina **%s**?", discipline.Name)

	ConfirmationButton(s, interaction, customIDConfirmation, customIDCancel, messageContent)
}

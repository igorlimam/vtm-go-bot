package view

import (
	"fmt"
	"log"
	"vtm-go-bot/model"

	"github.com/bwmarrin/discordgo"
)

func AddDisciplineView(s *discordgo.Session, interaction *discordgo.InteractionCreate, discipline *model.Discipline) {
	customID := "add-discipline-modal"
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

	var choices []*discordgo.ApplicationCommandOptionChoice
	for _, discipline := range disciplines {
		choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
			Name:  discipline.Name,
			Value: fmt.Sprintf("%d", discipline.ID),
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

func ShowDisciplineInfoView(s *discordgo.Session, interaction *discordgo.InteractionCreate, discipline model.Discipline) {
	embed := &discordgo.MessageEmbed{
		Title:       discipline.Name,
		Description: discipline.Description,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Tipo",
				Value:  discipline.Dtype,
				Inline: true,
			},
			{
				Name:   "Ressonância",
				Value:  discipline.Resonance,
				Inline: true,
			},
			{
				Name:   "Ameaça",
				Value:  discipline.Threat,
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

func ConfirmDeleteDiscipline(s *discordgo.Session, interaction *discordgo.InteractionCreate, discipline model.Discipline) {
	confirmBtn := discordgo.Button{
		CustomID: "confirm-delete-discipline|" + fmt.Sprintf("%d", discipline.ID),
		Label:    "Sim",
		Style:    discordgo.DangerButton,
	}
	cancelBtn := discordgo.Button{
		CustomID: "cancel-delete-discipline",
		Label:    "Não",
		Style:    discordgo.SecondaryButton,
	}

	err := s.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("Tem certeza que deseja deletar a disciplina **%s**?", discipline.Name),
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{Components: []discordgo.MessageComponent{confirmBtn, cancelBtn}},
			},
		},
	})

	if err != nil {
		log.Println("Error responding to confirm delete discipline interaction:", err)
	}
}

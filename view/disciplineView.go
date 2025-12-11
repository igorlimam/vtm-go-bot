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
								CustomID:    "discipline-name",
								Label:       "Nome da Disciplina",
								Style:       discordgo.TextInputShort,
								Value:       nameValue,
								Placeholder: "Insira o nome da disciplina",
								Required:    true,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID:    "discipline-type",
								Label:       "Tipo da Disciplina",
								Style:       discordgo.TextInputShort,
								Value:       typeValue,
								Placeholder: "Insira o tipo da disciplina",
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID:    "discipline-resonance",
								Label:       "Ressonância da Disciplina",
								Style:       discordgo.TextInputShort,
								Value:       resonanceValue,
								Placeholder: "Insira a ressonância da disciplina",
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID:    "discipline-description",
								Label:       "Descrição da Disciplina",
								Style:       discordgo.TextInputParagraph,
								Value:       descriptionValue,
								Placeholder: "Insira a descrição da disciplina",
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID:    "discipline-threat",
								Label:       "Ameaça da Disciplina",
								Style:       discordgo.TextInputParagraph,
								Value:       threatValue,
								Placeholder: "Insira a ameaça da disciplina",
							},
						},
					},
				},
			},
		},
	)
	if err != nil {
		log.Println("Error responding with modal:", err)
	}
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

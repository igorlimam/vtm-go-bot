package view

import (
	"fmt"
	"log"
	"vtm-go-bot/model"

	"github.com/bwmarrin/discordgo"
)

func AddDisciplineView(s *discordgo.Session, interaction *discordgo.InteractionCreate) {
	s.InteractionRespond(
		interaction.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseModal,
			Data: &discordgo.InteractionResponseData{
				CustomID: "add-discipline-modal",
				Title:    "Adicionar Nova Disciplina",
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID: "discipline-name",
								Label:    "Nome da Disciplina",
								Style:    discordgo.TextInputShort,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID: "discipline-type",
								Label:    "Tipo da Disciplina",
								Style:    discordgo.TextInputShort,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{

							discordgo.TextInput{
								CustomID: "discipline-resonance",
								Label:    "Ressonância da Disciplina",
								Style:    discordgo.TextInputShort,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID: "discipline-description",
								Label:    "Descrição da Disciplina",
								Style:    discordgo.TextInputParagraph,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID: "discipline-threat",
								Label:    "Ameaça da Disciplina",
								Style:    discordgo.TextInputParagraph,
							},
						},
					},
				},
			},
		},
	)
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

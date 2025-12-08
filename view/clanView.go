package view

import (
	"fmt"
	"log"
	"strings"
	"vtm-go-bot/model"

	"github.com/bwmarrin/discordgo"
)

func StringSelectClanDisciplines(s *discordgo.Session, interaction *discordgo.InteractionCreate, disciplines []model.Discipline) {
	choices := make([]discordgo.SelectMenuOption, 0)
	for _, discipline := range disciplines {
		choices = append(choices, discordgo.SelectMenuOption{
			Label: discipline.Name,
			Value: fmt.Sprintf("%d", discipline.ID),
		})
	}

	min := 1
	selectMenu := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			&discordgo.SelectMenu{
				CustomID:    "select-disciplines-for-clan",
				Placeholder: "Selecione pelo menos uma disciplinas",
				MinValues:   &min,
				MaxValues:   len(choices),
				Options:     choices,
			},
		},
	}

	err := s.InteractionRespond(
		interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content:    "Escolha as disciplinas para o novo clã:",
				Components: []discordgo.MessageComponent{selectMenu},
			},
		},
	)

	if err != nil {
		log.Println("Error responding to interaction:", err)
	}
}

func AddClanView(s *discordgo.Session, interaction *discordgo.InteractionCreate, disciplineIDs []string) {

	suffix := strings.Join(disciplineIDs, "-")

	err := s.InteractionRespond(
		interaction.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseModal,
			Data: &discordgo.InteractionResponseData{
				CustomID: "add-clan-modal|" + suffix,
				Title:    "Adicionar Novo Clã",
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID: "clan-name",
								Label:    "Nome do Clã",
								Style:    discordgo.TextInputShort,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID: "clan-description",
								Label:    "Descrição do Clã",
								Style:    discordgo.TextInputParagraph,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID: "clan-bane",
								Label:    "Fraqueza do Clã",
								Style:    discordgo.TextInputParagraph,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID: "clan-compulsion",
								Label:    "Compulsão do Clã",
								Style:    discordgo.TextInputParagraph,
							},
						},
					},
				},
			},
		},
	)

	if err != nil {
		log.Println("Error responding to interaction:", err)
	}
}

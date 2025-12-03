package view

import "github.com/bwmarrin/discordgo"

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

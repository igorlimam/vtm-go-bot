package view

import "github.com/bwmarrin/discordgo"

func AddDisciplineView(s *discordgo.Session, interaction *discordgo.InteractionCreate) {
	s.InteractionRespond(
		interaction.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseModal,
			Data: &discordgo.InteractionResponseData{
				CustomID: "add-discipline-modal",
				Title:    "Add New Discipline",
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID: "discipline-name",
								Label:    "Discipline Name",
								Style:    discordgo.TextInputShort,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID: "discipline-type",
								Label:    "Discipline Type",
								Style:    discordgo.TextInputShort,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{

							discordgo.TextInput{
								CustomID: "discipline-resonance",
								Label:    "Discipline Resonance",
								Style:    discordgo.TextInputShort,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID: "discipline-threat",
								Label:    "Discipline Threat",
								Style:    discordgo.TextInputParagraph,
							},
						},
					},
				},
			},
		},
	)
}

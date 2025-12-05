package view

import (
	"fmt"
	"vtm-go-bot/model"

	"github.com/bwmarrin/discordgo"
)

func PowerSelectDisciplineView(s *discordgo.Session, interaction *discordgo.InteractionCreate, disciplines []model.Discipline) error {
	choices := make([]discordgo.SelectMenuOption, 0)
	for _, discipline := range disciplines {
		choices = append(choices, discordgo.SelectMenuOption{
			Label: discipline.Name,
			Value: fmt.Sprintf("%d", discipline.ID),
		})
	}

	return s.InteractionRespond(
		interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Escolha a disciplina para o novo poder:",
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.SelectMenu{
								CustomID: "select-discipline-for-power",
								Options:  choices,
							},
						},
					},
				},
			},
		},
	)

}

func AddPowerView(s *discordgo.Session, interaction *discordgo.InteractionCreate, disciplineId string) {
	err := s.InteractionRespond(
		interaction.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseModal,
			Data: &discordgo.InteractionResponseData{
				CustomID: "add-power-modal|" + disciplineId,
				Title:    "Adicionar Novo Poder",
				Components: []discordgo.MessageComponent{
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID: "power-name",
								Label:    "Nome do Poder",
								Style:    discordgo.TextInputShort,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID: "power-description",
								Label:    "Descrição do Poder",
								Style:    discordgo.TextInputParagraph,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID: "power-dice-pool",
								Label:    "Parada de Dados",
								Style:    discordgo.TextInputParagraph,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{

							discordgo.TextInput{
								CustomID: "power-cost",
								Label:    "Custo do Poder",
								Style:    discordgo.TextInputShort,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID: "power-duration",
								Label:    "Duração do Poder",
								Style:    discordgo.TextInputShort,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID: "power-system",
								Label:    "Sistema",
								Style:    discordgo.TextInputParagraph,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID: "power-type",
								Label:    "Tipo do Poder",
								Style:    discordgo.TextInputShort,
							},
						},
					},
					discordgo.ActionsRow{
						Components: []discordgo.MessageComponent{
							discordgo.TextInput{
								CustomID: "power-level",
								Label:    "Nível do Poder (1-10)",
								Style:    discordgo.TextInputShort,
							},
						},
					},
				},
			},
		},
	)
	if err != nil {
		fmt.Println("Error responding with modal:", err)
	}
}

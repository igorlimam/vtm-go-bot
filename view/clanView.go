package view

import (
	"fmt"
	"log"
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

	min := 3
	selectMenu := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			&discordgo.SelectMenu{
				CustomID:    "select-disciplines-for-clan",
				Placeholder: "Selecione pelo menos três disciplinas",
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

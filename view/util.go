package view

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func ResolveResponse(s *discordgo.Session, interaction *discordgo.InteractionCreate, response string) {
	err := s.InteractionRespond(
		interaction.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: response,
			},
		},
	)
	if err != nil {
		log.Printf("Failed to respond to interaction: %v", err)
	}
}

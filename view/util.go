package view

import (
	"fmt"
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func Modal(s *discordgo.Session, interaction *discordgo.InteractionCreate, customID string, title string, fields []map[string]string) {

	var actionRows []discordgo.MessageComponent

	for _, field := range fields {
		var style discordgo.TextInputStyle
		if field["style"] == "short" {
			style = discordgo.TextInputShort
		} else {
			style = discordgo.TextInputParagraph
		}
		actionRow := discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.TextInput{
					CustomID: field["customID"],
					Label:    field["label"],
					Style:    style,
					Value:    field["value"],
				},
			},
		}
		actionRows = append(actionRows, actionRow)
	}

	err := s.InteractionRespond(
		interaction.Interaction,
		&discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseModal,
			Data: &discordgo.InteractionResponseData{
				CustomID:   customID,
				Title:      title,
				Components: actionRows,
			},
		},
	)

	if err != nil {
		log.Printf("Error responding to %s:\n", customID)
		log.Println("Error responding to MODAL interaction:", err)
	}
}

func SelectMenu(s *discordgo.Session, interaction *discordgo.InteractionCreate, options []map[string]string, customID string, placeholder string, contentPlaceholder string) {

	var choices []discordgo.SelectMenuOption
	for _, option := range options {
		def, _ := strconv.ParseBool(option["default"])
		val, _ := strconv.Atoi(option["value"])
		choices = append(choices, discordgo.SelectMenuOption{
			Label:   option["label"],
			Value:   fmt.Sprintf("%d", val),
			Default: def,
		})
	}

	min := 1
	selectMenu := discordgo.ActionsRow{
		Components: []discordgo.MessageComponent{
			&discordgo.SelectMenu{
				CustomID:    customID,
				Placeholder: placeholder,
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
				Content:    contentPlaceholder,
				Components: []discordgo.MessageComponent{selectMenu},
			},
		},
	)

	if err != nil {
		log.Printf("Error responding to %s:\n", customID)
		log.Println("Error responding to SELECT MENU interaction:", err)
	}
}

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

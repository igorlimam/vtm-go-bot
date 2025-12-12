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

func SelectMenu(s *discordgo.Session, interaction *discordgo.InteractionCreate, options []map[string]string, customID string, placeholder string, contentPlaceholder string, max int) {

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
				MaxValues:   max,
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

func AutoComplete(s *discordgo.Session, interaction *discordgo.InteractionCreate, options []map[string]string) {

	var choices []*discordgo.ApplicationCommandOptionChoice
	for _, option := range options {
		choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
			Name:  option["label"],
			Value: option["value"],
		})
	}

	err := s.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionApplicationCommandAutocompleteResult,
		Data: &discordgo.InteractionResponseData{
			Choices: choices,
		},
	})

	if err != nil {
		log.Println("Error responding to AUTOCOMPLETE interaction:", err)
	}
}

func EmbedMessage(s *discordgo.Session, interaction *discordgo.InteractionCreate, fields []map[string]string, title string, description string) {

	embedFields := []*discordgo.MessageEmbedField{}
	for _, field := range fields {
		embedFields = append(embedFields, &discordgo.MessageEmbedField{
			Name:   field["name"],
			Value:  field["value"],
			Inline: field["inline"] == "true",
		})
	}

	embed := &discordgo.MessageEmbed{
		Title:       title,
		Description: description,
		Fields:      embedFields,
	}

	err := s.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	})

	if err != nil {
		log.Println("Error responding to EMBED MESSAGE interaction:", err)
	}
}

func ConfirmationButton(s *discordgo.Session, interaction *discordgo.InteractionCreate, customIDConfirmation string, customIDCancel string, contentMessage string) {
	confirmBtn := discordgo.Button{
		CustomID: customIDConfirmation,
		Label:    "Sim",
		Style:    discordgo.SecondaryButton,
	}
	cancelBtn := discordgo.Button{
		CustomID: customIDCancel,
		Label:    "NÃO",
		Style:    discordgo.DangerButton,
	}

	err := s.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: contentMessage,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{Components: []discordgo.MessageComponent{confirmBtn, cancelBtn}},
			},
		},
	})

	if err != nil {
		log.Println("Error responding to CONFIRM BUTTON interaction:", err)
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
		log.Printf("Failed to respond to RESPONSE interaction: %v", err)
	}
}

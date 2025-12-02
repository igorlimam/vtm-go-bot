package main

import (
	"log"
	"vtm-go-bot/controller"
	"vtm-go-bot/view"

	"github.com/bwmarrin/discordgo"
)

func RegisterCommands(session *discordgo.Session) {

	adminCommands := map[string]string{
		"ping":           "AM I ALIVE?",
		"add-discipline": "Add a new discipline",
	}

	userCommands := map[string]string{
		// Add user commands here
	}

	commands := make(map[string]string)

	for name, description := range adminCommands {
		commands[name] = description
	}
	for name, description := range userCommands {
		commands[name] = description
	}

	for name, description := range commands {
		_, err := session.ApplicationCommandCreate(
			session.State.User.ID,
			"",
			&discordgo.ApplicationCommand{
				Name:        name,
				Description: description,
			},
		)
		if err != nil {
			log.Fatalf("Cannot create '%v' command: %v", name, err)
		}
	}

	log.Println("Commands registered successfully.")
}

func RegisterHandlers(s *discordgo.Session, interaction *discordgo.InteractionCreate) {

	if interaction.Type == discordgo.InteractionModalSubmit {
		switch interaction.ModalSubmitData().CustomID {
		case "add-discipline-modal":
			controller.AddDiscipline(s, interaction)
		}
	}

	if interaction.Type != discordgo.InteractionApplicationCommand {
		return
	}

	switch interaction.ApplicationCommandData().Name {
	case "ping":
		s.InteractionRespond(
			interaction.Interaction,
			&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Pong",
				},
			},
		)
	case "add-discipline":
		view.AddDisciplineView(s, interaction)
	}

}

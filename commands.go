package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func RegisterCommands(session *discordgo.Session) {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "ping",
			Description: "AM I ALIVE?",
		},
	}

	for _, command := range commands {
		_, err := session.ApplicationCommandCreate(
			session.State.User.ID,
			"",
			command,
		)
		if err != nil {
			log.Fatalf("Cannot create '%v' command: %v", command.Name, err)
		}
	}

	log.Println("Commands registered successfully.")
}

func PingCommand() string {
	log.Println("Ping command invoked")
	return "pong"
}

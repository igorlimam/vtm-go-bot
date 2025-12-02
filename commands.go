package main

import (
	"log"
	"vtm-go-bot/repository"

	"github.com/bwmarrin/discordgo"
)

func RegisterCommands(session *discordgo.Session) {

	commands := map[string]string{
		"ping":           "AM I ALIVE?",
		"add-discipline": "Add a new discipline",
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

func PingCommand() string {
	log.Println("Ping command invoked")
	db := repository.ConnectDuckDB()
	defer db.Close()
	repository.CheckDDL(db)
	return "pong"
}

func AddDiscipline(interaction *discordgo.InteractionCreate) map[string]string {
	log.Println("AddDiscipline command invoked")

	log.Printf("Interaction Data: %+v\n", interaction.Data)

	return map[string]string{
		"status": "Discipline added successfully",
	}
}

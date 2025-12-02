package main

import (
	"log"
	"vtm-go-bot/controller"
	"vtm-go-bot/view"

	"github.com/bwmarrin/discordgo"
)

func checkGuildOwner(s *discordgo.Session, interaction *discordgo.InteractionCreate) {
	var guild, _ = s.State.Guild(interaction.GuildID)
	log.Printf("Comparing Guild Owner ID: %s, with Interaction User ID: %s", guild.OwnerID, interaction.Member.User.ID)
	if interaction.Member.User.ID != guild.OwnerID {
		s.InteractionRespond(
			interaction.Interaction,
			&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Apenas o dono do servidor pode usar este comando.",
				},
			},
		)
	}
}

func RegisterCommands(session *discordgo.Session) {
	commands := map[string]string{
		"ping":           "AM I ALIVE?",
		"add-disciplina": "Adiciona uma nova disciplina",
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

	if interaction.Type != discordgo.InteractionApplicationCommand {
		return
	}

	if interaction.Type == discordgo.InteractionModalSubmit {
		switch interaction.ModalSubmitData().CustomID {
		case "add-disciplina-modal":
			controller.AddDiscipline(s, interaction)
		}
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
	case "add-disciplina":
		checkGuildOwner(s, interaction)
		view.AddDisciplineView(s, interaction)
	}

}

package main

import (
	"log"
	"strings"
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
		"add-poder":      "Adiciona um novo poder a uma disciplina existente",
		"add-clan":       "Adiciona um novo clã",
	}

	for name, description := range commands {
		_, err := session.ApplicationCommandCreate(
			session.State.User.ID,
			session.State.Guilds[0].ID,
			&discordgo.ApplicationCommand{
				Name:        name,
				Description: description,
			},
		)
		if err != nil {
			log.Fatalf("Cannot create '%v' command: %v", name, err)
		}
	}

	readCommands := map[string]string{
		"disciplina": "Fornece informações sobre uma disciplina específica",
	}
	for name, description := range readCommands {
		_, err := session.ApplicationCommandCreate(
			session.State.User.ID,
			session.State.Guilds[0].ID,
			&discordgo.ApplicationCommand{
				Name:        name,
				Description: description,
				Options: []*discordgo.ApplicationCommandOption{
					{
						Type:         discordgo.ApplicationCommandOptionString,
						Name:         name,
						Description:  description,
						Required:     true,
						Autocomplete: true,
					},
				},
			},
		)
		if err != nil {
			log.Fatalf("Cannot create '%v' command: %v", name, err)
		}
	}

	log.Println("Commands registered successfully.")
	controller.CheckDDLController()
}

func RegisterHandlers(s *discordgo.Session, interaction *discordgo.InteractionCreate) {

	if interaction.Type == discordgo.InteractionModalSubmit {
		switch strings.Split(interaction.ModalSubmitData().CustomID, "|")[0] {
		case "add-discipline-modal":
			controller.AddDiscipline(s, interaction)
		case "add-power-modal":
			controller.AddPower(s, interaction, strings.Split(interaction.ModalSubmitData().CustomID, "|")[1])
		case "add-clan-modal":
			controller.AddClan(s, interaction, strings.Split(interaction.ModalSubmitData().CustomID, "|")[1])
		}
	}

	if interaction.Type == discordgo.InteractionMessageComponent {
		data := interaction.MessageComponentData()
		switch data.CustomID {
		case "select-discipline-for-power":
			view.AddPowerView(s, interaction, data.Values[0])
		case "select-disciplines-for-clan":
			view.AddClanView(s, interaction, data.Values)
			log.Printf("Selected disciplines for clan: %v", data.Values)
		}
	}

	if interaction.Type == discordgo.InteractionApplicationCommandAutocomplete {
		for _, opt := range interaction.ApplicationCommandData().Options {
			if opt.Focused && opt.Name == "disciplina" {
				view.DisciplinaInfoView(s, interaction, controller.GetAllDisciplines())
				return
			}
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
	case "add-disciplina":
		checkGuildOwner(s, interaction)
		view.AddDisciplineView(s, interaction)
	case "add-poder":
		checkGuildOwner(s, interaction)
		view.PowerSelectDisciplineView(s, interaction, controller.GetAllDisciplines())
	case "add-clan":
		checkGuildOwner(s, interaction)
		view.StringSelectClanDisciplines(s, interaction, controller.GetAllDisciplines())
	case "disciplina":
		disciplinaID := interaction.ApplicationCommandData().Options[0].StringValue()
		view.ShowDisciplineInfoView(s, interaction, controller.GetDisciplineByID(disciplinaID))
	}

}

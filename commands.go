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

	selectionCommands := map[string][]map[string]string{
		"disciplina": {
			{"name": "disciplina", "description": "Fornece informações sobre uma disciplina específica"},
		},
		"clan": {
			{"name": "clan", "description": "Fornece informações sobre um clã específico"},
		},
		"poder": {
			{"name": "disciplina", "description": "Disciplina do poder"},
			{"name": "poder", "description": "Fornece informações sobre um poder específico"},
		},
		"update-disciplina": {
			{"name": "disciplina", "description": "Disciplina a ser atualizada"},
		},
		"update-clan": {
			{"name": "clan", "description": "Clã a ser atualizado"},
		},
		"update-poder": {
			{"name": "disciplina", "description": "Disciplina do poder"},
			{"name": "poder", "description": "Poder a ser atualizado"},
		},
		"delete-disciplina": {
			{"name": "disciplina", "description": "Disciplina a ser deletada"},
		},
	}

	var all []*discordgo.ApplicationCommand

	for cmd, desc := range commands {
		all = append(all, &discordgo.ApplicationCommand{
			Name:        cmd,
			Description: desc,
		})
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

	for cmd, cmdMap := range selectionCommands {
		var options []*discordgo.ApplicationCommandOption
		for _, params := range cmdMap {
			options = append(options, &discordgo.ApplicationCommandOption{
				Type:         discordgo.ApplicationCommandOptionString,
				Name:         params["name"],
				Description:  params["description"],
				Required:     true,
				Autocomplete: true,
			})
		}
		all = append(all, &discordgo.ApplicationCommand{
			Name:        cmd,
			Description: "Mostrar informações sobre " + cmd,
			Options:     options,
		})
	}

	_, err := session.ApplicationCommandBulkOverwrite(
		session.State.User.ID,
		session.State.Guilds[0].ID,
		all,
	)
	if err != nil {
		log.Fatalf("Cannot bulk overwrite commands: %v", err)
	}

	log.Println("Commands registered successfully.")
	controller.CheckDDLController()
}

func RegisterHandlers(s *discordgo.Session, interaction *discordgo.InteractionCreate) {

	if interaction.Type == discordgo.InteractionModalSubmit {
		customID := strings.Split(interaction.ModalSubmitData().CustomID, "|")[0]
		customData := strings.Split(interaction.ModalSubmitData().CustomID, "|")[1]
		log.Printf("CUSTOM ID: %s", customID)
		switch customID {
		case "add-discipline-modal":
			status := controller.AddDiscipline(s, interaction)
			view.ResolveResponse(s, interaction, status)
		case "update-discipline-modal":
			status := controller.UpdateDiscipline(s, interaction, customData)
			view.ResolveResponse(s, interaction, status)
		case "add-power-modal":
			status := controller.AddPower(s, interaction, customData)
			view.ResolveResponse(s, interaction, status)
		case "add-clan-modal":
			status := controller.AddClan(s, interaction, customData)
			view.ResolveResponse(s, interaction, status)
		case "update-clan-modal":
			status := controller.UpdateClan(s, interaction, customData)
			view.ResolveResponse(s, interaction, status)
		case "update-power-modal":
			powerID := strings.Split(interaction.ModalSubmitData().CustomID, "|")[1]
			disciplineID := strings.Split(interaction.ModalSubmitData().CustomID, "|")[2]
			status := controller.UpdatePower(s, interaction, powerID, disciplineID)
			view.ResolveResponse(s, interaction, status)
		}
	}

	if interaction.Type == discordgo.InteractionMessageComponent {
		data := interaction.MessageComponentData()
		customID := strings.Split(data.CustomID, "|")[0]
		switch customID {
		case "select-discipline-for-power":
			view.AddPowerView(s, interaction, data.Values[0], nil)
		case "select-disciplines-for-clan":
			clanID := strings.Split(data.CustomID, "|")[1]

			if clanID != "0" {
				clan := controller.GetClanByID(clanID)
				view.AddClanView(s, interaction, data.Values, &clan)
			} else {
				view.AddClanView(s, interaction, data.Values, nil)
			}

			log.Printf("Selected disciplines for clan: %v", data.Values)
		case "confirm-delete-discipline":
			disciplineID := strings.Split(data.CustomID, "|")[1]
			status := controller.DeleteDiscipline(s, interaction, disciplineID)
			view.ResolveResponse(s, interaction, status)
		default:
			status := "Interação Cancelada"
			view.ResolveResponse(s, interaction, status)
		}
	}

	if interaction.Type == discordgo.InteractionApplicationCommandAutocomplete {
		for _, opt := range interaction.ApplicationCommandData().Options {
			if opt.Focused && opt.Name == "disciplina" {
				log.Printf("Focused option for disciplina: %s", opt.StringValue())
				view.DisciplinaInfoView(s, interaction, controller.GetAllDisciplines())
				return
			}
			if opt.Focused && opt.Name == "clan" {
				view.ClanInfoView(s, interaction, controller.GetAllClans())
				return
			}
			if opt.Focused && opt.Name == "poder" {
				var query string
				var disciplineID string
				for _, option := range interaction.ApplicationCommandData().Options {
					if option.Name == "disciplina" {
						disciplineID = option.StringValue()
					}
					if option.Focused {
						query = strings.ToLower(option.StringValue())
					}
				}
				log.Printf("Autocomplete query for power: %s", disciplineID)
				disciplinePowers := controller.GetDisciplinePowersByID(disciplineID)
				view.PowerInfoView(s, interaction, query, disciplinePowers)
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
		view.AddDisciplineView(s, interaction, nil)
	case "add-poder":
		checkGuildOwner(s, interaction)
		view.PowerSelectDisciplineView(s, interaction, controller.GetAllDisciplines())
	case "add-clan":
		checkGuildOwner(s, interaction)
		view.StringSelectClanDisciplines(s, interaction, controller.GetAllDisciplines(), nil, "")
	case "disciplina":
		disciplinaID := interaction.ApplicationCommandData().Options[0].StringValue()
		view.ShowDisciplineInfoView(s, interaction, controller.GetDisciplineByID(disciplinaID))
	case "clan":
		clanID := interaction.ApplicationCommandData().Options[0].StringValue()
		disciplines := controller.GetClanDisciplinesById(clanID)
		view.ShowClanInfoView(s, interaction, controller.GetClanByID(clanID), disciplines)
	case "poder":
		power := controller.GetPowerById(interaction.ApplicationCommandData().Options[1].StringValue())
		view.ShowPowerInfoView(s, interaction, power)
	case "update-disciplina":
		checkGuildOwner(s, interaction)
		discipline := controller.GetDisciplineByID(interaction.ApplicationCommandData().Options[0].StringValue())
		view.AddDisciplineView(s, interaction, &discipline)
	case "update-clan":
		checkGuildOwner(s, interaction)
		clanID := interaction.ApplicationCommandData().Options[0].StringValue()
		selectedDisciplines := controller.GetClanDisciplinesById(clanID)
		disciplines := controller.GetAllDisciplines()
		view.StringSelectClanDisciplines(s, interaction, disciplines, selectedDisciplines, clanID)
	case "update-poder":
		checkGuildOwner(s, interaction)
		disciplineID := interaction.ApplicationCommandData().Options[0].StringValue()
		powerID := interaction.ApplicationCommandData().Options[1].StringValue()
		power := controller.GetPowerById(powerID)
		view.AddPowerView(s, interaction, disciplineID, &power)
	case "delete-disciplina":
		checkGuildOwner(s, interaction)
		disciplineID := interaction.ApplicationCommandData().Options[0].StringValue()
		view.ConfirmDeleteDiscipline(s, interaction, controller.GetDisciplineByID(disciplineID))
	default:
		status := "Interação Cancelada"
		view.ResolveResponse(s, interaction, status)
	}

}

package service

import "github.com/bwmarrin/discordgo"

func ModalToMap(interaction *discordgo.InteractionCreate) map[string]string {
	dataModal := interaction.ModalSubmitData().Components
	values := map[string]string{}

	for _, row := range dataModal {
		if row, ok := row.(*discordgo.ActionsRow); ok {
			for _, component := range row.Components {
				if input, ok := component.(*discordgo.TextInput); ok {
					values[input.CustomID] = input.Value
				}
			}
		}
	}

	return values
}

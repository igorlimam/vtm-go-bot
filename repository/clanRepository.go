package repository

import "vtm-go-bot/model"

func AddClan(clanName string, description string, bane string, compulsion string, desciplines []model.Discipline) map[string]string {
	clan := model.Clan{
		Name:        clanName,
		Description: description,
		Bane:        bane,
		Compulsion:  compulsion,
		Disciplines: desciplines,
	}
	InsertIntoTable(&clan)
	return map[string]string{"status": "Clan added successfully"}
}

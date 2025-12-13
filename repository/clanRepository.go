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

func UpdateClan(id uint, clanName string, description string, bane string, compulsion string, desciplines []model.Discipline) map[string]string {
	clan := model.Clan{
		ID:          id,
		Name:        clanName,
		Description: description,
		Bane:        bane,
		Compulsion:  compulsion,
		Disciplines: desciplines,
	}
	UpdateTable(&clan)
	DB.Model(&clan).Association("Disciplines").Clear()
	if len(desciplines) > 0 {
		DB.Model(&clan).Association("Disciplines").Append(desciplines)
	}
	return map[string]string{"status": "Clan updated successfully"}
}

func GetAllClans() []model.Clan {
	var clans []model.Clan
	GetAll(&clans)
	return clans
}

func GetClanByID(id uint) model.Clan {
	var clan model.Clan
	GetByID(&clan, id)
	return clan
}

func DeleteClan(id uint) map[string]string {
	var clan model.Clan
	GetByID(&clan, id)
	DeleteFromTable(&clan)
	return map[string]string{"status": "Clan deleted successfully"}
}

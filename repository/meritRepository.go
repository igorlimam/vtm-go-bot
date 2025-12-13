package repository

import (
	"vtm-go-bot/model"
)

func AddMerit(name string, description string, kind string, levelsInfo string) map[string]string {

	meritToBeInserted := model.Merit{
		Name:        name,
		Description: description,
		Kind:        kind,
		LevelsInfo:  levelsInfo,
	}
	InsertIntoTable(&meritToBeInserted)
	return map[string]string{"status": "Mérito adicionado com sucesso!"}
}

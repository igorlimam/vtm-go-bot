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

func GetMeritByID(id uint) model.Merit {
	var merit model.Merit
	GetByID(&merit, id)
	return merit
}

func GetMeritsByKind(kind string) []model.Merit {
	var merits []model.Merit
	GetByField(&merits, "kind", kind)
	return merits
}

func UpdateMerit(id uint, name string, description string, kind string, levelsInfo string) map[string]string {

	meritToBeUpdated := model.Merit{
		ID:          id,
		Name:        name,
		Description: description,
		Kind:        kind,
		LevelsInfo:  levelsInfo,
	}
	UpdateTable(&meritToBeUpdated)
	return map[string]string{"status": "Mérito atualizado com sucesso!"}
}

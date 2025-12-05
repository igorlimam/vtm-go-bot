package repository

import (
	"log"
	"vtm-go-bot/model"
)

func AddDiscipline(name string, dtype string, resonance string, threat string, description string) map[string]string {
	disciplineToBeInserted := model.Discipline{
		Name:        name,
		Dtype:       dtype,
		Resonance:   resonance,
		Threat:      threat,
		Description: description,
	}
	InsertIntoTable(&disciplineToBeInserted)
	log.Printf("Inserted Discipline: %s with ID: %d\n", name, disciplineToBeInserted.ID)
	return map[string]string{"status": "Disciplina adicionada com sucesso!"}
}

func GetAllDisciplines() []model.Discipline {
	var disciplines []model.Discipline
	DB.Find(&disciplines)
	return disciplines
}

func getDisciplineById(id uint) model.Discipline {
	var discipline model.Discipline
	DB.First(&discipline, id)
	return discipline
}

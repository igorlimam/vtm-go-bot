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

func UpdateDiscipline(id uint, name string, dtype string, resonance string, threat string, description string) map[string]string {
	disciplineToBeUpdated := model.Discipline{
		ID:          id,
		Name:        name,
		Dtype:       dtype,
		Resonance:   resonance,
		Threat:      threat,
		Description: description,
	}
	UpdateTable(&disciplineToBeUpdated)
	log.Printf("Updated Discipline ID %d: %s\n", id, name)
	return map[string]string{"status": "Disciplina atualizada com sucesso!"}
}

func GetAllDisciplines() []model.Discipline {
	var disciplines []model.Discipline
	GetAll(&disciplines)
	return disciplines
}

func GetDisciplineById(id uint) model.Discipline {
	var discipline model.Discipline
	GetByID(&discipline, id)
	return discipline
}

func GetDisciplineByName(name string) model.Discipline {
	var disciplines []model.Discipline
	GetByField(&disciplines, "name", name)
	return disciplines[0]
}

func DeleteDiscipline(id uint) map[string]string {
	var discipline model.Discipline
	GetByID(&discipline, id)
	if discipline.ID == 0 {
		log.Printf("Discipline with ID %d not found for deletion\n", id)
		return map[string]string{"status": "Disciplina não encontrada!"}
	}
	DeleteFromTable(&discipline)
	log.Printf("Deleted Discipline ID %d: %s\n", id, discipline.Name)
	return map[string]string{"status": "Disciplina deletada com sucesso!"}
}

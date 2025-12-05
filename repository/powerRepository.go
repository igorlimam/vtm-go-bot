package repository

import "vtm-go-bot/model"

func AddPower(disciplineId uint, name string, description string,
	dicePool string, cost string, duration string, system string,
	ptype string, level int) map[string]string {

	powerToBeInserted := model.Power{
		DisciplineID: disciplineId,
		Name:         name,
		Description:  description,
		DicePool:     dicePool,
		Cost:         cost,
		Duration:     duration,
		System:       system,
		Kind:         ptype,
		Level:        level,
	}
	InsertIntoTable(&powerToBeInserted)
	return map[string]string{"status": "Poder adicionado com sucesso!"}
}

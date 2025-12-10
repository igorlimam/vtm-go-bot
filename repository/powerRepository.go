package repository

import "vtm-go-bot/model"

func AddPower(disciplineId uint, name string, description string,
	dicePool string, cost string, duration string, system string,
	ptype string, amalgam string, level int) map[string]string {

	powerToBeInserted := model.Power{
		DisciplineID: disciplineId,
		Name:         name,
		Description:  description,
		DicePool:     dicePool,
		Cost:         cost,
		Duration:     duration,
		System:       system,
		Kind:         ptype,
		Amalgam:      amalgam,
		Level:        level,
	}
	InsertIntoTable(&powerToBeInserted)
	return map[string]string{"status": "Poder adicionado com sucesso!"}
}

func GetAllPowers() []model.Power {
	var powers []model.Power
	GetAll(&powers)
	return powers
}

func GetPowersByDiciplineId(disciplineId uint) []model.Power {
	var powers []model.Power
	GetByField(&powers, "discipline_id", disciplineId)
	return powers
}

func GetPowerById(id uint) model.Power {
	var power model.Power
	GetByID(&power, id)
	return power
}

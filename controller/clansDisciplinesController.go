package controller

import (
	"vtm-go-bot/model"
	"vtm-go-bot/service"
)

func GetClanDisciplinesById(clanId string) []model.Discipline {
	return service.GetClanDisciplinesByIdService(clanId)
}

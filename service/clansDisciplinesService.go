package service

import (
	"log"
	"strconv"
	"vtm-go-bot/model"
	"vtm-go-bot/repository"
)

func GetClanDisciplinesByIdService(clanId string) []model.Discipline {
	id, err := strconv.Atoi(clanId)
	if err != nil {
		log.Printf("Invalid clan ID: %v", err)
		return []model.Discipline{}
	}

	disciplines := repository.GetClanDisciplinesByIdRepository(uint(id))
	return disciplines
}

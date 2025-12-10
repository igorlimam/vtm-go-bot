package repository

import (
	"log"
	"vtm-go-bot/model"
)

func GetClanDisciplinesByIdRepository(clanId uint) []model.Discipline {
	log.Printf("Fetching disciplines for clan ID: %d", clanId)
	var clanDisciplines []model.ClansDisciplines

	GetByField(&clanDisciplines, "clan_id", clanId)

	var disciplines []model.Discipline
	for _, cd := range clanDisciplines {
		var discipline model.Discipline
		GetByID(&discipline, cd.DisciplineID)
		disciplines = append(disciplines, discipline)
	}

	return disciplines
}

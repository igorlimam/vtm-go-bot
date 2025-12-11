package model

type ClansDisciplines struct {
	ID           uint `gorm:"primaryKey;autoIncrement"`
	ClanID       uint `gorm:"not null;index"`
	DisciplineID uint `gorm:"not null;index"`
}

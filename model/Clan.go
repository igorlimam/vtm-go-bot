package model

type Clan struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Bane        string
	Compulsion  string
	Disciplines []Discipline `gorm:"many2many:clans_disciplines;"`
}

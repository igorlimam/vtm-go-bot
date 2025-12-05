package model

type Power struct {
	ID           uint `gorm:"primaryKey"`
	DisciplineID uint
	Discipline   Discipline `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Name        string
	Description string
	DicePool    string
	Cost        string
	Duration    string
	System      string
	Kind        string
	Level       int
}

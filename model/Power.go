package model

type Power struct {
	ID           uint       `gorm:"primaryKey"`
	DisciplineID uint       `gorm:"not null;index"`
	Discipline   Discipline `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Name        string
	Description string
	DicePool    string
	Cost        string
	Duration    string
	System      string
	Kind        string
	Amalgam     string `gorm:"default:-"`
	Level       int    `gorm:"index"`
}

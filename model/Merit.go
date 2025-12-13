package model

type Merit struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
	Kind        string `gorm:"index"`
	LevelsInfo  string
}

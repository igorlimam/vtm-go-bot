package model

type Discipline struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Dtype       string
	Resonance   string
	Threat      string
	Description string
}

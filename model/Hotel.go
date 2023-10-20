package model

type Hotel struct {
	ID       int `gorm:"primaryKey"`
	Nome     string
	Endereco string
}

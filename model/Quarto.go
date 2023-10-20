package model

import "bohotel/database"

type Quarto struct {
	ID              int `gorm:"primaryKey;column:QuartoID"`
	Tipo            string
	Preco           float64
	Disponibilidade bool
}

func ObterQuartoPorID(quartoID int) (Quarto, error) {
	var quarto Quarto
	db := database.Db
	result := db.Where("QuartoID  = ?", quartoID).First(&quarto)
	if result.Error != nil {
		return Quarto{}, result.Error
	}
	return quarto, nil
}

func ObterPrecoQuartoPorID(quartoID int) (float64, error) {
	var quarto Quarto
	db := database.Db
	result := db.Select("Preco").Where("QuartoID = ?", quartoID).First(&quarto)
	if result.Error != nil {
		return 0, result.Error
	}
	return quarto.Preco, nil
}

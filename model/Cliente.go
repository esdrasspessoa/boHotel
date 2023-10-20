package model

import "bohotel/database"

type Cliente struct {
	ID             int `gorm:"primaryKey"`
	Nome           string
	NumeroTelefone int       `gorm:"column:NumeroTelefone"`
	Reservas       []Reserva // Relação um-para-muitos com Reserva
}

func CriarCliente(cliente Cliente) (Cliente, error) {
	db := database.Db
	result := db.Create(&cliente)
	if result.Error != nil {
		return Cliente{}, result.Error
	}
	return cliente, nil
}

func ListarClientes() ([]Cliente, error) {
	var clientes []Cliente
	db := database.Db
	result := db.Find(&clientes)
	if result.Error != nil {
		return nil, result.Error
	}
	return clientes, nil
}

func ObterClientePorID(clienteID int) (Cliente, error) {
	var cliente Cliente
	db := database.Db
	result := db.Where("ID = ?", clienteID).First(&cliente)
	if result.Error != nil {
		return Cliente{}, result.Error
	}
	return cliente, nil
}

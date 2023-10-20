package service

import (
	"bohotel/database"
	"bohotel/model"
	"errors"
)

func CriarReserva(cliente model.Cliente, quarto model.Quarto, dataReserva string) (model.Reserva, error) {
	//Verifica a disponibilidade do quarto
	if !VerificarDisponibilidade(quarto, dataReserva) {
		return model.Reserva{}, errors.New("Quarto não disponível para data da reserva.")
	}

	//Cria a reserva
	reserva := model.Reserva{
		ClienteID:     cliente.ID,
		Quarto:        quarto,
		DataReserva:   dataReserva,
		PrecoTotal:    quarto.Preco,
		StatusReserva: "Pendente",
	}

	result := database.Db.Create(&reserva)
	if result.Error != nil {
		return model.Reserva{}, result.Error
	}

	return reserva, nil
}

func AtualizarReserva(reserva model.Reserva) error{
	result := database.Db.Save(&reserva)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func ExcluirReserva(id int) error {
	result := database.Db.Delete(&model.Reserva{}, id)
	if result.Error != nil{
		return result.Error
	}
	return nil
}

func VerificarDisponibilidade(quarto model.Quarto, dataReserva string) bool {
	var reservaExistente model.Reserva
	result := database.Db.Where("QuartoNumero = ? AND DataReserva = ?", quarto.ID, dataReserva).First(&reservaExistente)

	if result.Error == nil {
		// Se uma reserva existente for encontrada, o quarto não está disponível
		return false
	}

	// Se nenhuma reserva existente for encontrada, o quarto está disponível
	return true
}

func ListarReservas() ([]model.Reserva, error) {
    var reservas []model.Reserva
    result := database.Db.Preload("Cliente").Preload("Quarto").Find(&reservas)
    if result.Error != nil {
        return nil, result.Error
    }
    return reservas, nil
}
package main

import (
	"bohotel/database"
	"bohotel/model"
	"bohotel/service"
	"fmt"
	"time"
)

func main() {
	database.InitDB()

	for {
        fmt.Println("Selecione uma opção:")
        fmt.Println("1. Criar Cliente")
        fmt.Println("2. Listar Clientes")
        fmt.Println("3. Criar Reserva")
        fmt.Println("4. Listar Reservas")
        fmt.Println("5. Sair")

        var opcao int
        fmt.Scanln(&opcao)

        switch opcao {
        case 1:
            criarCliente()
        case 2:
            listarClientes()
        case 3:
            criarReserva()
        case 4:
            listarReservas()
        case 5:
            fmt.Println("Saindo...")
            return
        default:
            fmt.Println("Opção inválida. Tente novamente.")
        }
    }
	
}


func criarCliente() {
	var cliente model.Cliente
	fmt.Println("Nome do Cliente")
	fmt.Scanln(&cliente.Nome)
	fmt.Println("Numero de telefone:")
	fmt.Scanln(&cliente.NumeroTelefone)

	_, err := model.CriarCliente(cliente)
	if err != nil {
		fmt.Printf("Erro ao criar cliente: %v\n", err)
	}else {
		fmt.Println("Cliente criado com sucesso!")
	}
}

func listarClientes() {
	clientes, err := model.ListarClientes()
	if err != nil {
		fmt.Printf("Erro ao listar clientes: %v\n", err)
	} else {
		fmt.Println("Clientes:")
		for _, cliente := range clientes{
			fmt.Printf("ID: %d, Nome: %s, Número de Telefone: %d\n", cliente.ID, cliente.Nome, cliente.NumeroTelefone)
		}
	}
}

func criarReserva() {
	var reserva model.Reserva

	fmt.Println("ID do Cliente:")
	var clienteID int
	fmt.Scanln(&clienteID)

	cliente, err := model.ObterClientePorID(clienteID)
	if err != nil {
		fmt.Printf("Erro ao obter o cliente: %v\n", err)
		return
	}

	fmt.Println("ID do Quarto:")
	var quartoID int 
	fmt.Scanln(&quartoID)

	quarto, err := model.ObterQuartoPorID(quartoID)
	if err != nil {
		fmt.Printf("Erro ao obter o quarto: %v\n", err)
		return
	}

	fmt.Println("Data da Reserva (formato: DD/MM/AAAA):")
	var dataReservaInput string
	fmt.Scanln(&dataReservaInput)

	// Converter a data no formato "DD/MM/AAAA" para um objeto time.Time
	dataReserva, err := time.Parse("02/01/2006", dataReservaInput)
	if err != nil {
		fmt.Printf("Erro ao converter data: %v\n", err)
		return
	}

	// Converter a data para o formato "AAAA-MM-DD" (formato aceito pelo MySQL)
	dataReservaStr := dataReserva.Format("2006-01-02")

	// Obter o preço do quarto
	precoQuarto, err := model.ObterPrecoQuartoPorID(quartoID)
	if err != nil {
		fmt.Printf("Erro ao obter o preço do quarto: %v\n", err)
		return
	}

	if !service.VerificarDisponibilidade(quarto, dataReservaStr){
		fmt.Println("Quarto não disponivel para data da reserva!")
		return
	}

	reserva.ClienteID = cliente.ID
	reserva.Quarto = quarto
	reserva.DataReserva = dataReservaStr
	reserva.PrecoTotal = precoQuarto
	reserva.StatusReserva = "Pendente"

	_, err = service.CriarReserva(cliente, quarto, dataReservaStr)
	if err != nil {
		fmt.Printf("Erro ao criar reserva: %v\n", err)
	} else {
		fmt.Println("Reserva criada com sucesso!")
	}
}

func listarReservas() {
	reservas, err := service.ListarReservas()
	if err != nil {
		fmt.Printf("Erro ao listar reservas: %v\n", err)
	} else {
		fmt.Println("Reservas:")
		for _, reserva := range reservas {
			fmt.Printf("ID: %d, Cliente: %s, Quarto: %s, Data de Reserva: %s, Preço Total: %.2f, Status: %s\n", reserva.ID, reserva.Cliente.Nome, reserva.Quarto.Tipo, reserva.DataReserva, reserva.PrecoTotal, reserva.StatusReserva)
		}
	}
}
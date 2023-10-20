package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

/*
func InitDB() {
	var err error
	Db, err = sql.Open("mysql", "root:191103@tcp(127.0.0.1:3306)/hotel_db")
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		return
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println("Erro ao verificar conexão com o banco de dados:", err)
		return
	}
}
*/

func InitDB() {
	var err error

	// Configuração da conexão com o banco de dados
	dsn := "root:191103@tcp(127.0.0.1:3306)/hotel_db?parseTime=True"

	// Inicialização do GORM
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Erro ao conectar ao banco de dados: " + err.Error())
	}

	// Configuração do NamingStrategy
	Db.NamingStrategy = schema.NamingStrategy{
		TablePrefix:   "hotel_db.", // Use o nome do seu banco de dados como prefixo se necessário
		SingularTable: true,        // Use tabelas no singular
	}
}

package model

type Reserva struct {
	ID            int     `gorm:"primaryKey"`
	ClienteID     int
	Cliente       Cliente `gorm:"foreignKey:ClienteID"`
	QuartoNumero  int
	Quarto        Quarto  `gorm:"foreignKey:QuartoNumero"`
	DataReserva   string
	PrecoTotal    float64
	StatusReserva string
}

package models

import "time"

type Resultado struct {
	ID         int
	PacienteID int
	ExamenID   int
	Valor      string
	Fecha      time.Time
}

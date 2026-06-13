package models

import (
	"errors"
	"time"
)

type Resultado struct {
	id         int
	pacienteID int
	examenID   int
	valor      string
	fecha      time.Time
}

func (r *Resultado) Validar() error {
	if r.valor == "" {
		return errors.New("el valor del resultado no puede estar vacío")
	}
	return nil
}

func (r *Resultado) ObtenerValor() string   { return r.valor }
func (r *Resultado) ObtenerPacienteID() int { return r.pacienteID }
func (r *Resultado) ObtenerExamenID() int   { return r.examenID }

func NuevoResultado(pID, eID int, valor string) *Resultado {
	return &Resultado{
		pacienteID: pID,
		examenID:   eID,
		valor:      valor,
		fecha:      time.Now(),
	}
}

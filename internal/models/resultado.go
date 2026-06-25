package models

import (
	"errors"
	"time"
)

type Resultado struct {
	ID           int       `json:"id"`              // Cambiado a ID
	PacienteID   int       `json:"paciente_id"`     // Cambiado a PacienteID
	ExamenID     int       `json:"examen_id"`       // Cambiado a ExamenID
	Valor        string    `json:"valor_resultado"` // Cambiado a Valor
	Fecha        time.Time `json:"fecha"`           // Cambiado a Fecha
	ExamenNombre string    `json:"nombre_examen"`
}

func (r *Resultado) Validar() error {
	if r.Valor == "" {
		return errors.New("el valor del resultado no puede estar vacío")
	}
	return nil
}

func (r *Resultado) ObtenerValor() string   { return r.Valor }
func (r *Resultado) ObtenerPacienteID() int { return r.PacienteID }
func (r *Resultado) ObtenerExamenID() int   { return r.ExamenID }

func NuevoResultado(pID, eID int, valor string) *Resultado {
	return &Resultado{
		PacienteID:   pID, // Cambiado a mayúscula
		ExamenID:     eID, // Cambiado a mayúscula
		Valor:        valor,
		Fecha:        time.Now(),
		ExamenNombre: "Perfil Lipidico", // Asignamos un valor por defecto o lo pasas como parámetro
	}
}

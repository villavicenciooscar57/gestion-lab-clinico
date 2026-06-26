package models

import (
	"errors"
	"time"
)

type Resultado struct {
	ID           int       `json:"id"`
	PacienteID   int       `json:"paciente_id"`
	ExamenID     int       `json:"examen_id"`
	Valor        string    `json:"valor_resultado"`
	Fecha        time.Time `json:"fecha"`
	ExamenNombre string    `json:"nombre_examen"`
	Precio       float64   `json:"precio"`
}

// Validar verifica que los datos básicos estén presentes
func (r *Resultado) Validar() error {
	if r.PacienteID <= 0 {
		return errors.New("el ID del paciente es inválido")
	}
	if r.Valor == "" {
		return errors.New("el valor del resultado no puede estar vacío")
	}
	return nil
}

// Métodos auxiliares
func (r *Resultado) ObtenerValor() string   { return r.Valor }
func (r *Resultado) ObtenerPacienteID() int { return r.PacienteID }
func (r *Resultado) ObtenerExamenID() int   { return r.ExamenID }

// NuevoResultado crea una instancia lista para ser enviada a la BD
func NuevoResultado(pID, eID int, valor string, nombreExamen string) *Resultado {
	return &Resultado{
		PacienteID:   pID,
		ExamenID:     eID,
		Valor:        valor,
		Fecha:        time.Now(),
		ExamenNombre: nombreExamen,
	}
}

package models

import "errors"

type Examen struct {
	ID           int     `json:"id,omitempty"`
	NombreExamen string  `json:"nombre_examen"`
	Descripcion  string  `json:"descripcion"`
	Precio       float64 `json:"precio"`
}

// Validar verifica que los campos obligatorios del examen sean correctos
func (e *Examen) Validar() error {
	if e.NombreExamen == "" || e.Precio <= 0 {
		return errors.New("examen requiere nombre y precio > 0")
	}
	return nil
}

// NuevoExamen es el constructor para crear un examen de forma segura
func NuevoExamen(nombre, descripcion string, precio float64) *Examen {
	return &Examen{
		NombreExamen: nombre,
		Descripcion:  descripcion,
		Precio:       precio,
	}
}

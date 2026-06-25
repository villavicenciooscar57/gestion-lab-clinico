/* STREAMING_CHUNK:Añadiendo el método Validar dentro del paquete del modelo */
package models

import "errors"

type Paciente struct {
	ID       int    `json:"id" db:"id"`
	Nombre   string `json:"nombre" db:"nombre"`
	Apellido string `json:"apellido" db:"apellido"`
	Cedula   string `json:"cedula" db:"cedula"`
	Email    string `json:"email" db:"email"`
	Telefono string `json:"telefono" db:"telefono"`
}

func (p *Paciente) Validar() error {
	if p.Cedula == "" || p.Nombre == "" {
		return errors.New("paciente requiere cédula y nombre")
	}
	return nil
}

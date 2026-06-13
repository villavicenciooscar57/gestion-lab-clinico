/* STREAMING_CHUNK:Añadiendo el método Validar dentro del paquete del modelo */
package models

import "errors"

type Paciente struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Cedula   string `json:"cedula"`
	Email    string `json:"email"`
	Telefono string `json:"telefono"`
}

func (p *Paciente) Validar() error {
	if p.Cedula == "" || p.Nombre == "" {
		return errors.New("paciente requiere cédula y nombre")
	}
	return nil
}

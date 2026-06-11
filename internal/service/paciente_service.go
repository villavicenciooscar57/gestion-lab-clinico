package service

import (
	"errors"

	"github.com/villavicenciooscar57/gestion-lab-clinico/internal/models"
)

// PacienteService maneja la lógica de negocio de los pacientes
type PacienteService struct{}

// RegistrarPaciente valida que el paciente tenga datos mínimos antes de guardarlo
func (s *PacienteService) RegistrarPaciente(p models.Paciente) (models.Paciente, error) {
	if p.Nombre == "" {
		return models.Paciente{}, errors.New("el nombre es obligatorio")
	}
	if p.Cedula == "" {
		return models.Paciente{}, errors.New("la cédula es obligatoria")
	}

	// Aquí es donde más adelante conectaremos con la base de datos
	return p, nil
}

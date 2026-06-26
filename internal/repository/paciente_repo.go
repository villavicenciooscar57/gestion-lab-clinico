package repository

import (
	"fmt"

	"github.com/villavicenciooscar57/gestion-lab-clinico/internal/models"
)

// ObtenerPacientes trae todos los pacientes desde Supabase
func ObtenerPacientes() ([]models.Paciente, error) {
	var pacientes []models.Paciente
	_, err := ClienteSupabase.From("pacientes").Select("*", "exact", false).ExecuteTo(&pacientes)
	return pacientes, err
}

// InsertarPaciente  incluye los exámenes solicitados
func InsertarPaciente(p models.Paciente) error {
	var pCreado []models.Paciente

	// Incluimos examenes_solicitados
	data := map[string]interface{}{
		"nombre":               p.Nombre,
		"apellido":             p.Apellido,
		"cedula":               p.Cedula,
		"email":                p.Email,
		"telefono":             p.Telefono,
		"examenes_solicitados": p.ExamenesSolicitados,
	}

	fmt.Printf("Enviando a Supabase: %+v\n", data)

	_, err := ClienteSupabase.From("pacientes").
		Insert(data, false, "", "", "").
		ExecuteTo(&pCreado)

	if err != nil {
		fmt.Printf("ERROR REAL DESDE SUPABASE: %v\n", err)
		return err
	}

	fmt.Printf("Paciente insertado con éxito, ID generado: %d\n", pCreado[0].ID)
	return nil
}

// BuscarPacientePorCedula busca un paciente específico mediante su cédula
func BuscarPacientePorCedula(cedula string) (*models.Paciente, error) {
	var pacientes []models.Paciente

	_, err := ClienteSupabase.From("pacientes").
		Select("*", "exact", false).
		Eq("cedula", cedula).
		ExecuteTo(&pacientes)

	if err != nil {
		return nil, fmt.Errorf("error al buscar paciente: %v", err)
	}

	if len(pacientes) == 0 {
		return nil, fmt.Errorf("paciente no encontrado")
	}

	return &pacientes[0], nil
}

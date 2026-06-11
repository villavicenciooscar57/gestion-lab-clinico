package repository

import (
	"fmt"

	"github.com/villavicenciooscar57/gestion-lab-clinico/internal/models"
)

// ObtenerPacientes trae todos los pacientes desde Supabase
func ObtenerPacientes() ([]models.Paciente, error) {
	var pacientes []models.Paciente
	// Se añade el guion bajo (_) para capturar el valor de retorno que ignoramos
	_, err := ClienteSupabase.From("pacientes").Select("*", "exact", false).ExecuteTo(&pacientes)
	return pacientes, err
}

// InsertarPaciente inserta un paciente y muestra la respuesta detallada
func InsertarPaciente(p models.Paciente) error {
	data, _, err := ClienteSupabase.From("pacientes").Insert(p, false, "", "", "").Execute()

	if err != nil {
		return fmt.Errorf("error al ejecutar la inserción: %v", err)
	}

	fmt.Printf("DEBUG: Respuesta cruda del servidor: %s\n", string(data))

	return nil
}

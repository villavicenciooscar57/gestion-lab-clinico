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
	var pCreado []models.Paciente

	// 1. Ejecutar el insert
	_, err := ClienteSupabase.From("pacientes").
		Insert(p, false, "", "", "").
		ExecuteTo(&pCreado)

	if err != nil {
		return err
	}

	// Devolvemos el paciente con su ID real generado
	fmt.Printf("Paciente insertado con éxito, ID asignado: %d\n", pCreado[0].ID)

	return nil
}

// BuscarPacientePorCedula busca un paciente específico mediante su cédula
func BuscarPacientePorCedula(cedula string) (*models.Paciente, error) {
	var pacientes []models.Paciente

	// Ejecutamos la consulta en Supabase filtrando por cédula
	// El método Eq filtra la columna 'cedula' por el valor recibido
	_, err := ClienteSupabase.From("pacientes").
		Select("*", "exact", false).
		Eq("cedula", cedula).
		ExecuteTo(&pacientes)

	if err != nil {
		return nil, fmt.Errorf("error al buscar paciente por cédula: %v", err)
	}

	// Si no encontramos registros, retornamos un error claro
	if len(pacientes) == 0 {
		return nil, fmt.Errorf("paciente no encontrado con cédula: %s", cedula)
	}

	// Retornamos el primer paciente encontrado
	return &pacientes[0], nil
}

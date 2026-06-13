package repository

import (
	"encoding/json" // Necesario para leer los datos de la base
	"fmt"

	"github.com/villavicenciooscar57/gestion-lab-clinico/internal/models"
)

// InsertarExamen envía un nuevo tipo de examen a la tabla 'examenes'
func InsertarExamen(e *models.Examen) error {
	if err := e.Validar(); err != nil {
		return fmt.Errorf("error de validación: %w", err)
	}

	_, _, err := ClienteSupabase.From("examenes").Insert(e, false, "", "", "").Execute()
	if err != nil {
		return fmt.Errorf("error al insertar examen: %v", err)
	}
	return nil
}

// ObtenerExamenes consulta la lista de exámenes disponibles
func ObtenerExamenes() ([]models.Examen, error) {
	resp, _, err := ClienteSupabase.From("examenes").Select("*", "", false).Execute()
	if err != nil {
		return nil, err
	}

	var examenes []models.Examen
	err = json.Unmarshal([]byte(resp), &examenes)
	if err != nil {
		return nil, err
	}
	return examenes, nil
}

package repository

import (
	"encoding/json"
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

// ObtenerExamenPorNombre consulta un examen específico por su nombre
func ObtenerExamenPorNombre(nombre string) (models.Examen, error) {
	var examenes []models.Examen

	_, err := ClienteSupabase.From("examenes").
		Select("*", "exact", false).
		Eq("nombre_examen", nombre).
		ExecuteTo(&examenes)

	if err != nil {
		return models.Examen{}, fmt.Errorf("error al conectar con base de datos: %v", err)
	}

	if len(examenes) == 0 {
		return models.Examen{}, fmt.Errorf("examen no encontrado: %s", nombre)
	}

	return examenes[0], nil
}

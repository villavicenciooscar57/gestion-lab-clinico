package repository

import (
	"fmt"

	"github.com/villavicenciooscar57/gestion-lab-clinico/internal/models"
)

// ObtenerExamenes trae todos los exámenes registrados
func ObtenerExamenes() ([]models.Examen, error) {
	var examenes []models.Examen
	_, err := ClienteSupabase.From("examenes").Select("*", "exact", false).ExecuteTo(&examenes)
	return examenes, err
}

// InsertarExamen agrega un nuevo tipo de examen al catálogo
func InsertarExamen(e models.Examen) error {
	data, _, err := ClienteSupabase.From("examenes").Insert(e, false, "", "", "").Execute()
	if err != nil {
		return fmt.Errorf("error al insertar examen: %v", err)
	}

	fmt.Printf("DEBUG: Respuesta cruda del servidor: %s\n", string(data))
	return nil
}

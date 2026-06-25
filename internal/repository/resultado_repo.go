package repository

import (
	"fmt"

	"github.com/villavicenciooscar57/gestion-lab-clinico/internal/models"
)

func InsertarResultado(r *models.Resultado) error {
	// Validamos usando el método que vive en models/resultado.go
	if err := r.Validar(); err != nil {
		return fmt.Errorf("error de validación: %w", err)
	}

	resultadoMap := map[string]interface{}{
		"paciente_id":     r.ObtenerPacienteID(),
		"examen_id":       r.ObtenerExamenID(),
		"valor_resultado": r.ObtenerValor(),
	}

	_, _, err := ClienteSupabase.From("resultados").Insert(resultadoMap, false, "", "", "").Execute()
	return err
}

// ObtenerResultadosPorPacienteID busca todos los resultados asociados a un paciente
func ObtenerResultadosPorPacienteID(pacienteID int) ([]models.Resultado, error) {
	var resultados []models.Resultado

	// Filtramos en la tabla 'resultados' donde 'paciente_id' coincida
	_, err := ClienteSupabase.From("resultados").
		Select("*", "exact", false).
		Eq("paciente_id", fmt.Sprintf("%d", pacienteID)).
		ExecuteTo(&resultados)

	if err != nil {
		return nil, fmt.Errorf("error al obtener resultados: %v", err)
	}

	return resultados, nil
}

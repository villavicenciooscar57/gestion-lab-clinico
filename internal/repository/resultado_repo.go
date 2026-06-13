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

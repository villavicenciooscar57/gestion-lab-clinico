package repository

import (
	"fmt"
	"time"

	"github.com/villavicenciooscar57/gestion-lab-clinico/internal/models"
)

// InsertarResultado guarda un nuevo examen en la tabla 'resultados'
func InsertarResultado(r *models.Resultado) error {
	if r.Fecha.IsZero() {
		r.Fecha = time.Now()
	}

	resultadoMap := map[string]interface{}{
		"paciente_id":     r.PacienteID,
		"examen_id":       r.ExamenID,
		"valor_resultado": r.Valor,
		"fecha":           r.Fecha,
		"nombre_examen":   r.ExamenNombre,
		"precio":          r.Precio,
	}

	_, _, err := ClienteSupabase.From("resultados").
		Insert(resultadoMap, false, "", "", "").
		Execute()

	if err != nil {
		return fmt.Errorf("error al insertar resultado: %v", err)
	}

	return nil
}

// ObtenerResultadosPorPacienteID busca resultados históricos de un paciente
func ObtenerResultadosPorPacienteID(pacienteID int) ([]models.Resultado, error) {
	var resultados []models.Resultado
	_, err := ClienteSupabase.From("resultados").
		Select("*", "exact", false).
		Eq("paciente_id", fmt.Sprintf("%d", pacienteID)).
		ExecuteTo(&resultados)

	return resultados, err
}

// ObtenerPrecioExamen consulta el precio de un examen desde la tabla 'examenes'
func ObtenerPrecioExamen(examenID int) (float64, error) {
	var examen struct {
		Precio float64 `json:"precio"`
	}
	_, err := ClienteSupabase.From("examenes").
		Select("precio", "exact", false).
		Eq("id", fmt.Sprintf("%d", examenID)).
		Single().
		ExecuteTo(&examen)

	return examen.Precio, err
}

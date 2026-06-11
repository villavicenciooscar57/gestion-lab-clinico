package repository

import (
	"github.com/villavicenciooscar57/gestion-lab-clinico/internal/models"
)

type ResultadoDetallado struct {
	ID               int
	Valor            string
	PacienteNombre   string
	PacienteApellido string
	ExamenNombre     string
}

func InsertarResultado(r models.Resultado) error {
	resultadoMap := map[string]interface{}{
		"paciente_id":     r.PacienteID,
		"examen_id":       r.ExamenID,
		"valor_resultado": r.Valor,
	}

	_, _, err := ClienteSupabase.From("resultados").Insert(resultadoMap, false, "", "", "").Execute()
	return err
}

func ObtenerResultadosDetallados() ([]ResultadoDetallado, error) {
	var rawData []map[string]interface{}

	query := "id, valor_resultado, pacientes(nombre, apellido), examenes(nombre_examen)"
	_, err := ClienteSupabase.From("resultados").Select(query, "exact", false).ExecuteTo(&rawData)
	if err != nil {
		return nil, err
	}

	var resultados []ResultadoDetallado
	for _, item := range rawData {
		paciente := item["pacientes"].(map[string]interface{})
		examen := item["examenes"].(map[string]interface{})

		resultados = append(resultados, ResultadoDetallado{
			ID:               int(item["id"].(float64)),
			Valor:            item["valor_resultado"].(string),
			PacienteNombre:   paciente["nombre"].(string),
			PacienteApellido: paciente["apellido"].(string),
			ExamenNombre:     examen["nombre_examen"].(string),
		})
	}
	return resultados, nil
}

// Buscar resulatdos.//
func BuscarResultadosPorFiltro(nombre string, fecha string) ([]ResultadoDetallado, error) {
	var rawData []map[string]interface{}

	query := ClienteSupabase.From("resultados").
		Select("id, valor_resultado, pacientes(nombre, apellido), examenes(nombre_examen)", "exact", false)

	query = query.Filter("pacientes.nombre", "ilike", "%"+nombre+"%")

	if fecha != "" {
		query = query.Filter("created_at", "gte", fecha).Filter("created_at", "lt", fecha+"T23:59:59")
	}

	_, err := query.ExecuteTo(&rawData)
	if err != nil {
		return nil, err
	}

	var resultados []ResultadoDetallado
	for _, item := range rawData {
		paciente := item["pacientes"].(map[string]interface{})
		examen := item["examenes"].(map[string]interface{})

		resultados = append(resultados, ResultadoDetallado{
			ID:               int(item["id"].(float64)),
			Valor:            item["valor_resultado"].(string),
			PacienteNombre:   paciente["nombre"].(string),
			PacienteApellido: paciente["apellido"].(string),
			ExamenNombre:     examen["nombre_examen"].(string),
		})
	}
	return resultados, nil
}

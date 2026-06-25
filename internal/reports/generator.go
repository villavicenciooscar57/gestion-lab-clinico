package reports

import (
	"html/template"
	"os"

	"github.com/villavicenciooscar57/gestion-lab-clinico/internal/models"
)

// ReporteData es una estructura de datos que combina Paciente y Resultado
type ReporteData struct {
	Paciente  *models.Paciente
	Resultado *models.Resultado
}

func GenerarReporteHTML(p *models.Paciente, r *models.Resultado) error {
	tmpl, err := template.ParseFiles("internal/reports/template.html")
	if err != nil {
		return err
	}

	fileName := "reporte_" + p.Cedula + ".html"
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	data := ReporteData{Paciente: p, Resultado: r}
	return tmpl.Execute(file, data)
}

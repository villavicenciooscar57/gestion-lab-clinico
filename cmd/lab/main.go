package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/villavicenciooscar57/gestion-lab-clinico/internal/models"
	"github.com/villavicenciooscar57/gestion-lab-clinico/internal/repository"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SUPABASE_KEY")

	if url == "" || key == "" {
		log.Fatal("Las variables SUPABASE_URL o SUPABASE_KEY están vacías")
	}

	err = repository.InicializarSupabase(url, key)
	if err != nil {
		log.Fatalf("Error inicializando Supabase: %v", err)
	}

	fmt.Println("--- Iniciando sistema integral de laboratorio ---")

	// Lógica de Pacientes
	nuevoPaciente := models.Paciente{
		Nombre:   "Oscar",
		Apellido: "Villavicencio",
		Cedula:   "1712345678",
		Email:    "oscar@uide.edu.ec",
		Telefono: "0998877665",
	}
	err = repository.InsertarPaciente(nuevoPaciente)
	if err != nil {
		log.Printf("Aviso (Pacientes): %v", err)
	}

	// Lógica de Exámenes
	nuevoExamen := models.Examen{
		NombreExamen: "Hemograma Completo",
		Descripcion:  "Análisis de sangre básico y detallado",
		Precio:       15.50,
	}
	err = repository.InsertarExamen(nuevoExamen)
	if err != nil {
		log.Printf("Aviso (Exámenes): %v", err)
	}

	// Lógica de Resultados
	fmt.Println("--- Probando Módulo de Resultados (Join) ---")
	nuevoResultado := models.Resultado{
		PacienteID: 1,
		ExamenID:   1,
		Valor:      "120 mg/dL",
	}

	err = repository.InsertarResultado(nuevoResultado)
	if err != nil {
		log.Printf("Error al insertar resultado: %v", err)
	}

	// Consultar resultados
	resultados, err := repository.ObtenerResultadosDetallados()
	if err != nil {
		log.Printf("Error al obtener resultados: %v", err)
	} else {
		for _, r := range resultados {
			// Usamos los nombres correctos: r.PacienteNombre, r.PacienteApellido, r.ExamenNombre
			fmt.Printf("ID: %d | Paciente: %s %s | Examen: %s | Valor: %s\n",
				r.ID, r.PacienteNombre, r.PacienteApellido, r.ExamenNombre, r.Valor)
		}
	}
	// --- Probando Búsqueda ---
	fmt.Println("\n--- Buscando resultados de 'Oscar' ---")
	resultadosBusqueda, err := repository.BuscarResultadosPorFiltro("Oscar", "")
	if err != nil {
		log.Printf("Error al buscar: %v", err)
	} else {
		for _, r := range resultadosBusqueda {
			fmt.Printf("Encontrado -> ID: %d | Paciente: %s %s | Valor: %s\n",
				r.ID, r.PacienteNombre, r.PacienteApellido, r.Valor)
		}
	}
}

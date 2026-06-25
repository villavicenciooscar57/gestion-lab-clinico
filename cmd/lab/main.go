package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/villavicenciooscar57/gestion-lab-clinico/internal/models"
	"github.com/villavicenciooscar57/gestion-lab-clinico/internal/reports"
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
		log.Fatal("Variables de entorno vacías")
	}

	err = repository.InicializarSupabase(url, key)
	if err != nil {
		log.Fatalf("Error inicializando Supabase: %v", err)
	}

	fmt.Println("--- Iniciando sistema integral de laboratorio ---")

	fmt.Println("\n--- Catálogo de Exámenes Disponibles ---")
	examenes, err := repository.ObtenerExamenes()
	if err != nil {
		log.Printf("Error al obtener catálogo: %v", err)
	} else {
		for _, ex := range examenes {
			fmt.Printf("ID: %d | Examen: %-20s | Precio: $%.2f\n", ex.ID, ex.NombreExamen, ex.Precio)
		}
	}
	fmt.Println("--------------------------------------\n")

	nuevoPaciente := models.Paciente{
		Nombre:   "Carol",
		Apellido: "Villacis",
		Cedula:   "1705008642",
		Email:    "carolvi@uide.edu.ec",
		Telefono: "0987812222",
	}
	err = repository.InsertarPaciente(nuevoPaciente)
	if err != nil {
		log.Printf("Aviso (Pacientes): %v", err)
	}
	// 3. LÓGICA DE RESULTADOS
	fmt.Println("--- Probando Módulo de Resultados ---")

	pReal, err := repository.BuscarPacientePorCedula("1705008642")

	if err != nil {
		log.Printf("Error: No se pudo localizar al paciente para el resultado: %v", err)
	} else {
		fmt.Printf("ID real obtenido de BD: %d\n", pReal.ID)

		idExamenDeseado := 41

		// Verificamos si el examen existe en la lista que ya obtuviste arriba
		examenExiste := false
		for _, ex := range examenes {
			if ex.ID == idExamenDeseado {
				examenExiste = true
				break
			}
		}

		if examenExiste {
			nuevoResultado := models.NuevoResultado(pReal.ID, idExamenDeseado, "155 mg/dL")
			err = repository.InsertarResultado(nuevoResultado)
			if err != nil {
				log.Printf("Error al insertar resultado: %v", err)
			} else {
				fmt.Println("¡Resultado insertado exitosamente!")
				miPaciente := *pReal // Usamos el paciente real con su ID correcto
				miResultado := nuevoResultado
				err = reports.GenerarReporteHTML(&miPaciente, miResultado)
				if err != nil {
					log.Printf("Error al generar reporte: %v", err)
				} else {
					fmt.Println("Reporte generado con éxito: reporte_" + miPaciente.Cedula + ".html")
				}
			}
		} else {
			log.Printf("Error: El examen con ID %d no existe. No se insertó nada.", idExamenDeseado)
		}
	}
	fmt.Println("--- Probando Módulo de Exámenes ---")
	nuevoExamen := models.NuevoExamen("Perfil Cardiaco", "Análisis de enzimas y función Cardiaca", 58.00)
	err = repository.InsertarExamen(nuevoExamen)
	if err != nil {
		log.Printf("Error al insertar examen: %v", err)
	} else {
		fmt.Println("¡Examen insertado exitosamente en la base de datos!")
	}

	fmt.Print("Ingrese la cédula a buscar: ")
	var cedulaInput string
	fmt.Scanln(&cedulaInput)

	paciente, err := repository.BuscarPacientePorCedula(cedulaInput)
	if err != nil {
		log.Printf("Aviso: %v", err)
	} else {
		fmt.Printf("Paciente encontrado: %s %s (ID: %d)\n", paciente.Nombre, paciente.Apellido, paciente.ID)
		resultados, err := repository.ObtenerResultadosPorPacienteID(paciente.ID)
		if err != nil {
			log.Printf("Aviso: No se pudieron obtener resultados: %v", err)
		} else {
			fmt.Println("\n--- Historial de Exámenes ---")
			if len(resultados) == 0 {
				fmt.Println("No hay exámenes registrados para este paciente.")
			} else {
				for _, res := range resultados {
					fmt.Printf("- Examen ID: %d | Resultado: %s\n", res.ExamenID, res.Valor)
				}
			}
		}
	}
}

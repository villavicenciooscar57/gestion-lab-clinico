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
	// Carga variables de entorno (.env) para conectarse a la nube
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando el archivo .env")
	}

	url := os.Getenv("SUPABASE_URL")
	key := os.Getenv("SUPABASE_KEY")

	// Valida que las credenciales no estén vacías
	if url == "" || key == "" {
		log.Fatal("Variables de entorno vacías")
	}

	// Inicializa la conexión con Supabase usando el repositorio central
	err = repository.InicializarSupabase(url, key)
	if err != nil {
		log.Fatalf("Error inicializando Supabase: %v", err)
	}

	fmt.Println("--- Iniciando sistema integral de laboratorio ---")

	// 1. LÓGICA DE CATÁLOGO: Obtiene y muestra los exámenes disponibles
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

	// 2. LÓGICA DE PACIENTES: Crea un objeto paciente y lo guarda en la BD
	nuevoPaciente := models.Paciente{
		Nombre:   "Juan",
		Apellido: "Perez",
		Cedula:   "1721494941",
		Email:    "garcia@uide.edu.ec",
		Telefono: "0879477675",
	}
	err = repository.InsertarPaciente(nuevoPaciente)
	if err != nil {
		log.Printf("Aviso (Pacientes): %v", err)
	}

	// 3. LÓGICA DE RESULTADOS: Valida que el examen exista antes de insertar
	fmt.Println("--- Probando Módulo de Resultados ---")
	idExamenDeseado := 23 // ID a probar basado en el catálogo impreso arriba

	// Bucle para buscar si el ID ingresado está en la lista de la BD
	encontrado := false
	for _, ex := range examenes {
		if ex.ID == 23 {
			encontrado = true
			break
		}
	}

	// Ejecuta la inserción solo si la validación es positiva
	if encontrado {
		fmt.Printf("Validación exitosa: El examen ID %d existe. Procediendo a insertar...\n", idExamenDeseado)
		nuevoResultado := models.NuevoResultado(34, idExamenDeseado, "120 mg/dL")
		err = repository.InsertarResultado(nuevoResultado)
		if err != nil {
			log.Printf("Error al insertar resultado: %v", err)
		} else {
			fmt.Println("¡Resultado insertado exitosamente!")
		}
	} else {
		log.Printf("Error: El examen con ID %d no existe. No se insertó nada.", idExamenDeseado)
	}

	// 4. LÓGICA DE EXÁMENES: Registra un nuevo tipo de prueba en el catálogo
	fmt.Println("--- Probando Módulo de Exámenes ---")
	nuevoExamen := models.NuevoExamen("Perfil Lipidico", "Análisis de enzimas y función Lipidica", 35.00)
	err = repository.InsertarExamen(nuevoExamen)
	if err != nil {
		log.Printf("Error al insertar examen: %v", err)
	} else {
		fmt.Println("¡Examen insertado exitosamente en la base de datos!")
	}
}

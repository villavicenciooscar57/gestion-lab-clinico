package main

// 1. IMPORTACIÓN DE PAQUETES Y DEPENDENCIAS DEL SISTEMA
// ============================================================================
// Acción: Declara las librerías nativas de Go y los paquetes internos del
// proyecto necesarios para el funcionamiento del servidor HTTP, serialización y base de datos.
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/villavicenciooscar57/gestion-lab-clinico/internal/models"
	"github.com/villavicenciooscar57/gestion-lab-clinico/internal/repository"
)

// ============================================================================
// 2. FUNCIÓN PRINCIPAL / ENRUTAMIENTO BASE: http://localhost:8080/
// ============================================================================
// Acción: Inicializa el entorno, conecta la base de datos cloud, registra las
// rutas de la API REST y levanta el servidor web en el puerto 8080.
func main() {
	// Carga de variables de entorno del archivo .env (Credenciales secretas)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando .env")
	}
	// Inicialización y enlace con la base de datos Supabase Cloud (PostgreSQL)
	err = repository.InicializarSupabase(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"))
	if err != nil {
		log.Fatalf("Error inicializando Supabase: %v", err)
	}
	// REGISTRO DE ENDPOINTS (SERVICIOS WEB DE LA API REST)
	http.HandleFunc("/api/paciente", pacienteHandler)
	http.HandleFunc("/api/resultados", resultadoHandler)
	http.HandleFunc("/api/resultados/paciente", listarResultadosHandler)
	http.HandleFunc("/api/examenes", listarExamenesHandler)
	// RUTA DE ACCESO WEB (ESTÁTICOS): http://localhost:8080/ (Página Principal)
	// Acción: Sirve la interfaz gráfica de usuario desde el directorio local './frontend'
	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	// Lanzamiento y escucha activa del servidor web
	fmt.Println("Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// ============================================================================
// 3. ENRUTADOR PRINCIPAL / MULTIPLEXOR DE RECURSO: `/api/paciente`
// ============================================================================
// Acción: Actúa como un switch de control. Recibe todas las peticiones a '/api/paciente'
// y las redirige a la función correspondiente según el método HTTP (GET o POST).
func pacienteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		buscarPacienteHandler(w, r)
	case http.MethodPost:
		crearPacienteHandler(w, r)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

// ============================================================================
// 4. RUTA DE ACCESO WEB: `GET http://localhost:8080/api/paciente?cedula=...`
// ============================================================================
// Acción: Recupera el parámetro 'cedula' enviado desde el frontend, consulta en
// Supabase y devuelve los datos demográficos del paciente serializados en JSON.
func buscarPacienteHandler(w http.ResponseWriter, r *http.Request) {
	// Extracción del parámetro query de la URL (?cedula=...)
	cedula := r.URL.Query().Get("cedula")
	paciente, err := repository.BuscarPacientePorCedula(cedula)
	if err != nil {
		http.Error(w, "No encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paciente)
}

// ============================================================================
// 5. RUTA DE ACCESO WEB: `POST http://localhost:8080/api/paciente`
// ============================================================================
// Acción: Deserializa el JSON enviado por el frontend con los datos del formulario,
// los mapea al Struct 'Paciente' y los inserta en la base de datos remota.
func crearPacienteHandler(w http.ResponseWriter, r *http.Request) {
	var p models.Paciente
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Error en los datos: "+err.Error(), http.StatusBadRequest)
		return
	}
	// Inserción en la tabla de base de datos
	if err := repository.InsertarPaciente(p); err != nil {
		http.Error(w, "Error interno: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Registrado"})
}

// ============================================================================
// 6. RUTA DE ACCESO WEB: `POST http://localhost:8080/api/resultados`
// ============================================================================
// Acción: Recibe los valores medidos en los exámenes. Primero valida que el examen
// exista en el catálogo de Supabase para obtener su ID/precio, construye el registro
// enlazándolo al ID del paciente y lo guarda de forma definitiva.
func resultadoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Solo permitido POST", http.StatusMethodNotAllowed)
		return
	}
	// Estructura temporal para capturar el cuerpo de la petición
	var req struct {
		PacienteID   int    `json:"paciente_id"`
		Valor        string `json:"valor"`
		NombreExamen string `json:"nombre_examen"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	// Buscamos el examen para obtener su ID y Precio
	examen, err := repository.ObtenerExamenPorNombre(req.NombreExamen)
	if err != nil {
		http.Error(w, "Examen no configurado", http.StatusNotFound)
		return
	}

	// VINCULACIÓN RELACIONAL: Mapeo y empaquetado final de la entidad Resultado
	res := models.Resultado{
		PacienteID:   req.PacienteID,
		ExamenID:     examen.ID,
		ExamenNombre: req.NombreExamen,
		Valor:        req.Valor,
		Precio:       examen.Precio,
	}
	// Persistencia final en la base de datos en la nube
	if err := repository.InsertarResultado(&res); err != nil {
		http.Error(w, "Error interno al guardar: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"mensaje": "Examen registrado"})
}

// ============================================================================
// 7. RUTA DE ACCESO WEB: `GET http://localhost:8080/api/resultados/paciente?paciente_id=...`
// ============================================================================
// Acción: Captura el ID del paciente, convierte la variable de texto a un entero,
// extrae todos los resultados analíticos históricos asociados a ese ID y los devuelve al frontend.
func listarResultadosHandler(w http.ResponseWriter, r *http.Request) {
	pID := r.URL.Query().Get("paciente_id")
	id, err := strconv.Atoi(pID)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	resultados, err := repository.ObtenerResultadosPorPacienteID(id)
	if err != nil {
		http.Error(w, "Error al obtener historial", http.StatusInternalServerError)
		return
	}
	// Serialización y despacho del JSON de historial clínico
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultados)
}

// ============================================================================
// 8. RUTA DE ACCESO WEB: `GET http://localhost:8080/api/examenes`
// ============================================================================
// Acción: Consulta la tabla maestra de exámenes y retorna la lista completa de
// pruebas configuradas para poblar de forma dinámica la interfaz del Frontend.
func listarExamenesHandler(w http.ResponseWriter, r *http.Request) {
	// Consulta directa al catálogo general en Supabase
	examenes, err := repository.ObtenerExamenes()
	if err != nil {
		http.Error(w, "Error al cargar catálogo", http.StatusInternalServerError)
		return
	}
	// Envío del arreglo serializado con los exámenes disponibles
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(examenes)
}

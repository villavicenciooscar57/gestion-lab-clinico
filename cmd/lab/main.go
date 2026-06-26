package main

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

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error cargando .env")
	}

	err = repository.InicializarSupabase(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"))
	if err != nil {
		log.Fatalf("Error inicializando Supabase: %v", err)
	}

	http.HandleFunc("/api/paciente", pacienteHandler)
	http.HandleFunc("/api/resultados", resultadoHandler)
	http.HandleFunc("/api/resultados/paciente", listarResultadosHandler)
	http.HandleFunc("/api/examenes", listarExamenesHandler)

	http.Handle("/", http.FileServer(http.Dir("./frontend")))

	fmt.Println("Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

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

func buscarPacienteHandler(w http.ResponseWriter, r *http.Request) {
	cedula := r.URL.Query().Get("cedula")
	paciente, err := repository.BuscarPacientePorCedula(cedula)
	if err != nil {
		http.Error(w, "No encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(paciente)
}

func crearPacienteHandler(w http.ResponseWriter, r *http.Request) {
	var p models.Paciente
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Error en los datos: "+err.Error(), http.StatusBadRequest)
		return
	}
	if err := repository.InsertarPaciente(p); err != nil {
		http.Error(w, "Error interno: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"mensaje": "Registrado"})
}

func resultadoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Solo permitido POST", http.StatusMethodNotAllowed)
		return
	}

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

	// Mapeo al modelo
	res := models.Resultado{
		PacienteID:   req.PacienteID,
		ExamenID:     examen.ID,
		ExamenNombre: req.NombreExamen,
		Valor:        req.Valor,
		Precio:       examen.Precio,
	}

	if err := repository.InsertarResultado(&res); err != nil {
		http.Error(w, "Error interno al guardar: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"mensaje": "Examen registrado"})
}

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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resultados)
}

func listarExamenesHandler(w http.ResponseWriter, r *http.Request) {
	examenes, err := repository.ObtenerExamenes()
	if err != nil {
		http.Error(w, "Error al cargar catálogo", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(examenes)
}

package models

type Paciente struct {
	ID       int    `json:"id,omitempty"`
	Nombre   string `json:"nombre"`   // Cambia aquí si tu columna tiene mayúscula
	Apellido string `json:"apellido"` // Cambia aquí si tu columna tiene mayúscula
	Cedula   string `json:"cedula"`   // Cambia aquí si tu columna tiene mayúscula
	Email    string `json:"email"`    // Cambia aquí si tu columna tiene mayúscula
	Telefono string `json:"telefono"` // Cambia aquí si tu columna tiene mayúscula
}

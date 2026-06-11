package models

type Examen struct {
	ID           int     `json:"id,omitempty"`
	NombreExamen string  `json:"nombre_examen"`
	Descripcion  string  `json:"descripcion"`
	Precio       float64 `json:"precio"`
}

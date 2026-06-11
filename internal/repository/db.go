package repository

import (
	"github.com/supabase-community/supabase-go"
)

// ClienteSupabase almacena la instancia del cliente para ser usada en todo el repositorio.
var ClienteSupabase *supabase.Client

// InicializarSupabase configura el cliente oficial utilizando la URL y la llave pública.
// Al usar esta conexión vía HTTPS, evitamos los bloqueos de puerto 5432 en entornos restringidos.
func InicializarSupabase(url, key string) error {
	// Opciones nil por defecto; el cliente gestionará la comunicación segura vía HTTPS
	client, err := supabase.NewClient(url, key, nil)
	if err != nil {
		return err
	}
	ClienteSupabase = client
	return nil
}

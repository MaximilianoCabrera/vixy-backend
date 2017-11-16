package models

import "time"

type Comentario struct {
	ID        int
	Usuario   Usuario
	contenido string
	fecha     time.Time
}

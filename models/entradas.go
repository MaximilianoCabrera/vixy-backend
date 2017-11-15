package models

import "time"

type Entrada struct {
	ID         int
	Titulo     string
	Subtitulo  string
	Contenido  string
	Topic      int
	Fecha      time.Time
	Valoracion float32
	IDCiudad   int
	Visitas    int
	IDimagen   int
	Url        string
}

package models

import "time"

type Entrada struct {
	ID         int
	Titulo     string
	Subtitulo  string
	Contenido  string
	Topic      Topic
	Fecha      time.Time
	Valoracion float32
	IDCiudad   Ciudad
	Visitas    int
	Imagen     Imagen
	Url        string
}

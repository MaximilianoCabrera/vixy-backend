package models

type Usuario struct {
	ID          int
	Nombre      string
	Apellido    string
	Nick        string
	Email       string
	Pass	    string
	TipoUsuario int
	IDImagen	int
}

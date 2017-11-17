package models

type Usuario struct {
	ID          int		`json: "id"`
	Nombre      string	`json: "nombre"`
	Apellido    string	`json: "apellido"`
	Nick        string	`json: "nick"`
	Email       string	`json: "email"`
	Password    string	`json: "password"`
	TipoUsuario	int		`json: "idTipoUsuario"`
	Imagen		int		`json: "idImagen"`
}

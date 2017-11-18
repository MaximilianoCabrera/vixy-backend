package models

type GlobalModel struct {
	User        Usuario
	Imagen      Imagen
	TipoUsuario TipoUsuario
}

type GlobalModels struct {
	User        []Usuario
	Imagen      []Imagen
	TipoUsuario []TipoUsuario
}

type Users struct {
	Usuario []Usuario
}
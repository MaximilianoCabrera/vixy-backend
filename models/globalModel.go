package models

type GlobalModel struct {
	User        Usuario
	Imagen      Imagen
	TipoUsuario TipoUsuario
	Pais		Pais
	Continente  Continente
}

type GlobalModels struct {
	User        []Usuario
	Imagen      []Imagen
	TipoUsuario []TipoUsuario
	Pais		[]Pais
	Continente  []Continente
}
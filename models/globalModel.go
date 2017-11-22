package models

type GlobalModel struct {
	Actividad         Actividad
	Ciudad            Ciudad
	Comentario        Comentario
	ComentarioEntrada ComentarioEntrada
	Continente        Continente
	Entrada           Entrada
	Imagen            Imagen
	ImagenCiudad      ImagenCiudad
	ImagenEntrada     ImagenEntrada
	ImagenPais        ImagenPais
	Pais              Pais
	TipoUsuario       TipoUsuario
	Topic             Topic
	User              Usuario
	UsuarioEntrada    UsuarioEntrada
}

type GlobalModels struct {
	Actividad         []Actividad
	Ciudad            []Ciudad
	Comentario        []Comentario
	ComentarioEntrada []ComentarioEntrada
	Continente        []Continente
	Entrada           []Entrada
	Imagen            []Imagen
	ImagenCiudad      []ImagenCiudad
	ImagenEntrada     []ImagenEntrada
	ImagenPais        []ImagenPais
	Pais              []Pais
	TipoUsuario       []TipoUsuario
	Topic             []Topic
	User              []Usuario
	UsuarioEntrada    []UsuarioEntrada
}

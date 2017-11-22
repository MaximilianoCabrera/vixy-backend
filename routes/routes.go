package routes

import (
	"../actions"
	"github.com/gorilla/mux"
	"net/http"
)

func NewMainRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	rutas := []Routes{
		routes,
		//actividad,
		ciudad,
		//comentario,
		//comentarioEntrada,
		continente,
		//entrada,
		//imagen,
		//imagenCiudad,
		//imagenEntrada,
		//imagenPais,
		pais,
		//tipoUsuario,
		//topic,
		usuario,
		//usuarioEntrada,

	}

	for _, ruts := range rutas{
		for _, route := range ruts {
			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				HandlerFunc(route.HandleFunc)
		}
	}
	return router
}

//Route - Modelo de ruta
type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

//Routes - Listado de Route
type Routes []Route

var routes = Routes{
	//Index
	Route{
		"Index",
		"GET",
		"/",
		actions.Index,
	},
	Route{
		"Cerrar App",
		"GET",
		"/salir",
		actions.Close,
	},
}



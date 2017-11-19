package routes

import (
	"../actions"
	"github.com/gorilla/mux"
	"net/http"
)

//NewMainRouter creo las rutas
func NewMainRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			HandlerFunc(route.HandleFunc)
	}

	for _, route := range user {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			HandlerFunc(route.HandleFunc)
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
	//EntradasGetAll
	Route{
		"EntradasGetAll",
		"GET",
		"/entradas",
		actions.EntradasGetAll,
	},
}



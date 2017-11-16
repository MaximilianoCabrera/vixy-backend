package routes

import (
	"../actions"
	"github.com/gorilla/mux"
	"net/http"
)

//NewMainRouter creo las rutas
func NewMainRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
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
	//UsuarioCreate
	Route{
		"UserCreate",
		"POST",
		"/usuario",
		actions.UsuarioCreate,
	},
	//UserGetAll
	Route{
		"UserGetAll",
		"GET",
		"/usuarios",
		actions.UserGetAll,
	},
	//UserGetOne
	Route{
		"UserGetOne",
		"GET",
		"/usuarios/{params}",
		actions.UserGetOne,
	},
	//EntradasGetAll
	Route{
		"EntradasGetAll",
		"GET",
		"/entradas",
		actions.EntradasGetAll,
	},
}

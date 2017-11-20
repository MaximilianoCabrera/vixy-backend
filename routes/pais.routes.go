package routes

import (
	"../actions"
)

var pais = Routes{
	//PaisCreate
	Route{
		"PaisCreate",
		"POST",
		"/pais",
		actions.PaisCreate,
	},
	//PaisGet
	Route{
		"PaisGet",
		"GET",
		"/pais",
		actions.PaisGet,
	},
	//PaisGetBy
	Route{
		"PaisGetBy",
		"GET",
		"/pais/{id}",
		actions.PaisGetByID,
	},
	//PaisUpdate
	Route{
		"PaisUpdate",
		"PUT",
		"/pais/update",
		actions.PaisUpdate,
	},
	//PaisDelete
	Route{
		"PaisDelete",
		"DELETE",
		"/pais/{id}",
		actions.PaisDelete,
	},
}
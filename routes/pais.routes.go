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
}
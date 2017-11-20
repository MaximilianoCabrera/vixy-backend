package routes

import (
	"../actions"
)

var continente = Routes{
	//ContinenteCreate
	Route{
		"ContinenteCreate",
		"POST",
		"/continente",
		actions.ContinenteCreate,
	},
	//ContinenteGet
	Route{
		"ContinenteGet",
		"GET",
		"/continente",
		actions.ContinenteGet,
	},
/*
	//ContinenteGetBy
	Route{
		"ContinenteGetBy",
		"GET",
		"/continente/{id}",
		actions.ContinenteGetByID,
	},
	//ContinenteUpdate
	Route{
		"ContinenteUpdate",
		"PUT",
		"/continente/update/{id}",
		actions.ContinenteUpdate,
	},
	//ContinenteDelete
	Route{
		"ContinenteDelete",
		"DELETE",
		"/continente/{id}",
		actions.ContinenteDelete,
	},
*/
}
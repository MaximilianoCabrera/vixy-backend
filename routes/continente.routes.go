package routes

import (
	"../actions"
)

var continente = Routes{
	//Create
	Route{
		"ContinenteCreate",
		"POST",
		"/continente",
		actions.ContinenteCreate,
	},
	//Get (GetAll - GetBy)
	Route{
		"ContinenteGet",
		"GET",
		"/continente",
		actions.ContinenteGetBy,
	},
	//GetByID
	Route{
		"ContinenteGetBy",
		"GET",
		"/continente/{id}",
		actions.ContinenteGetByID,
	},
	//Update
	Route{
		"ContinenteUpdate",
		"PUT",
		"/continente/update/{id}",
		actions.ContinenteUpdate,
	},
	//Delete
	Route{
		"ContinenteDelete",
		"DELETE",
		"/continente/{id}",
		actions.ContinenteDelete,
	},
}
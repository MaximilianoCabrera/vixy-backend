package routes

import "../actions"

var user = Routes{
	//UserCreate
	Route{
		"UserCreate",
		"POST",
		"/usuario",
		actions.UserCreate,
	},
	//UserGetAll
	Route{
		"UserGet",
		"GET",
		"/usuario",
		actions.UserGet,
	},

	//UserGetBy
	Route{
		"UserGetBy",
		"GET",
		"/usuario/{id}",
		actions.UserGetByID,
	},
	//UserUpdate
	Route{
		"UserUpdate",
		"PUT",
		"/usuario/update/{id}",
		actions.UserUpdate,
	},
	//UserDelete
	Route{
		"UserDelete",
		"DELETE",
		"/usuario/{id}",
		actions.UserDelete,
	},
}
package routes

import "../actions"

var usuario = Routes{
	//UserCreate
	Route{
		"UserCreate",
		"POST",
		"/usuario",
		actions.UserCreate,
	},
	//UserGet
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
		"/usuario/update",
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

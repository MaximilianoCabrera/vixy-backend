package actions

import (
	"fmt"
	"net/http"
	"../models"
	"encoding/json"
	"../utilities"
	"log"
	"../dao/factory"
)

//response Usuario
func responseUser(w http.ResponseWriter, status int, results models.User){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(results)
}
//response Usuarios
func responseUsers(w http.ResponseWriter, status int, results []models.User, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err != nil{
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(results)
}

func UsersGetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo desde golang")
	var users []models.User
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}

	userDAO := factory.FactoryDAO(config.Engine)

	users, err = userDAO.GetAll()
	if err != nil{
		responseUsers(w, 404, nil, err)
	}
	responseUsers(w, 200, users, nil)
}
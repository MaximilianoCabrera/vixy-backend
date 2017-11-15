package actions

import (
	"net/http"
	"log"

	"../models"
	"../utilities"
	"../dao/factory"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
)

// response Usuario
func responseUser(w http.ResponseWriter, status int, results models.User, err error){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err != nil{
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(results)
}
// response Usuarios
func responseUsers(w http.ResponseWriter, status int, results []models.User, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err != nil{
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(results)
}

func UserGetAll(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}
	userDAO := factory.UserFactoryDAO(config.Engine)

	users, err = userDAO.GetAll()
	if err != nil{
		responseUsers(w, 404, nil, err)
	}
	responseUsers(w, 200, users, nil)
}

func UserGetOne(w http.ResponseWriter, r *http.Request){

	vars  := mux.Vars(r)
	params := vars["id"]

	id,err := strconv.Atoi(params)
	if err != nil {
		log.Fatalln(err)
	}

	var user models.User
	config, err := utilities.GetConfiguration()
	if err != nil{
		log.Fatalln(err)
	}
	userDAO := factory.UserFactoryDAO(config.Engine)

	user, err = userDAO.GetByID(id)
	if err != nil{
		fmt.Println("ERROR")
		//responseUser(w, 404, user, err)
	}
	responseUser(w, 200, user, nil)
}

func UsuarioCreate(w http.ResponseWriter, r *http.Request){

	config, err := utilities.GetConfiguration()
	if err != nil{
		log.Fatalln(err)
	}
	userDAO := factory.UsuarioFactoryDAO(config.Engine)

	user := models.Usuario{}

	fmt.Print("Nombre: ")
	fmt.Scan(&user.Nombre)
	fmt.Print("Apellido: ")
	fmt.Scan(&user.Apellido)
	fmt.Print("Nick: ")
	fmt.Scan(&user.Nick)
	fmt.Print("Email: ")
	fmt.Scan(&user.Email)
	fmt.Print("Password: ")
	fmt.Scan(&user.Password)
	user.IDTipoUsuario = 2

	fmt.Println(user)

	err = userDAO.Create(&user)
	if err != nil{
		fmt.Print("Error: ", err)
		log.Fatalln()
	}

	log.Println("Usuario " + user.Nombre + " cargado correctamente")
}
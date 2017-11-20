package actions

import (
	"net/http"
	"log"
	"fmt"
	"encoding/json"

	"../models"
	"../utilities"
	"../dao/interfaces"
	"../dao/factory"
)

// response uno
func response(w http.ResponseWriter, status int, results models.GlobalModel, model string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	switch model {
	case "user":
		json.NewEncoder(w).Encode(results.User)
	case "imagen":
		json.NewEncoder(w).Encode(results.Imagen)
	case "continente":
		json.NewEncoder(w).Encode(results.Continente)
	case "pais":
		json.NewEncoder(w).Encode(results.Pais)
	}
}
// response varios
func responses(w http.ResponseWriter, status int, results models.GlobalModels, model string, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	switch model {
	case "user":
		json.NewEncoder(w).Encode(results.User)
	case "imagen":
		json.NewEncoder(w).Encode(results.Imagen)
	case "continente":
		json.NewEncoder(w).Encode(results.Continente)
	case "pais":
		json.NewEncoder(w).Encode(results.Pais)
	}
}

// Func config
func config() models.Configuration {
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}
	return config
}
// Func GlobalDAO
func globalDAO() (x interfaces.GlobalDAO) {
	x = factory.GlobalFactoryDAO(config().Engine)
	return x
}

// Func Errores
func msjError(err error) {
	if err != nil {
		fmt.Println("Error: " , err)
	}
}
func checkErr(model string, accion string, err error) {
	if err != nil {
		switch model{
		case "user":
			log.Println("No se pudo " + accion + " el Usuario. Error: ", err)
		case "imagen":
			log.Println("No se pudo " + accion + " la Imagen. Error: ", err)
		case "pais":
			log.Println("No se pudo " + accion + " el Pais. Error: ", err)
		case "continente":
			log.Println("No se pudo " + accion + " el Continente. Error: ", err)
		default:
			log.Println("Error: ", err)
		}
	}
}

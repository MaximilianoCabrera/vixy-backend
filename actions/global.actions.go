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
// Func Errores
func msjError(err error) {
	fmt.Println("Error: ", err)
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
// Func respuesta error
func checkErr(msj string, err error) {
	if err != nil {
		log.Println(msj, err)
	}
}

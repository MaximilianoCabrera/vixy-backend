package actions

import (
	"net/http"
	"log"

	"../models"
	"../utilities"
	"../dao/factory"
	"encoding/json"
)

// response Usuario
func responseEntrada(w http.ResponseWriter, status int, results models.Entrada, err error){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err != nil{
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(results)
}

// response Usuarios
func responseEntradas(w http.ResponseWriter, status int, results []models.Entrada, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err != nil{
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(results)
}

func EntradasGetAll(w http.ResponseWriter, r *http.Request) {
	var entradas []models.Entrada
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}

	entradasDAO := factory.EntradaFactoryDAO(config.Engine)

	entradas, err = entradasDAO.GetAll()
	if err != nil{
		responseEntradas(w, 404, nil, err)
	}
	responseEntradas(w, 200, entradas, nil)
}


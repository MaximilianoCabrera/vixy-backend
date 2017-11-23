package actions

import (
	"net/http"

	"../models"
	"fmt"
)

//Controlar//
func ContinenteCreate(w http.ResponseWriter, r *http.Request) {
	var err error
	x := models.GlobalModel{}

	a := r.URL.Query()
	x.Continente.Nombre = a.Get("continente")

	if  x.Continente.Nombre != ""{

		globalDAO := globalDAO()

		x, err = globalDAO.Create(&x, "continente")
		if err != nil {
			fmt.Println("No se pudo crear el continente.", err)
			response(w, 404, models.GlobalModel{Pais: x.Pais}, "continente", err)
		}
		response(w, 200, models.GlobalModel{Pais: x.Pais}, "continente", err)
	} else {
		fmt.Println("Por favor vuelta a intentar cargar los datos en otro momento.")
		fmt.Println("")
	}
}
func ContinenteGetBy(w http.ResponseWriter, r *http.Request) {
	var x models.GlobalModel

	a := r.URL.Query()
	x.Continente.Nombre = a.Get("nombre")

	if x.Continente.Nombre != ""{
		globalDAO := globalDAO()
		a, err := globalDAO.GetBy(x, "continente")
		if err != nil {
			responses(w, 404, a, "continente", err)
		}
		responses(w, 200, a, "continente", nil)
	} else {
		var x models.GlobalModels

		globalDAO := globalDAO()
		x, err := globalDAO.GetAll("continente")
		if err != nil {
			responses(w, 404, x, "continente", err)
		}
		responses(w, 200, x, "continente", nil)
	}
}
func ContinenteGetByID(w http.ResponseWriter, r *http.Request){

}
func ContinenteUpdate(w http.ResponseWriter, r *http.Request){}
func ContinenteDelete(w http.ResponseWriter, r *http.Request){}

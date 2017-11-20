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
	x.Pais.Nombre = a.Get("nombre")
	x.Pais.Continente = a.Get("continente")
	x.Pais.Moneda = a.Get("moneda")
	x.Pais.Usohorario = a.Get("usoHorario")
	x.Pais.Idioma = a.Get("idioma")

	if x.Pais.Nombre != "" &&
		x.Pais.Continente != "" &&
		x.Pais.Moneda != "" &&
		x.Pais.Usohorario != "" &&
		x.Pais.Idioma != "" {


		x.Continente.Nombre = x.Pais.Continente
		globalDAO := globalDAO()

		x, err = globalDAO.GetOne(x, "continente")
		checkErr("No se pudo obtener el continente.", err)

		x, err = globalDAO.Create(&x, "pais")
		if err != nil {
			fmt.Println("No se pudo crear el pais.", err)
			response(w, 404, models.GlobalModel{Pais: x.Pais}, "pais", err)
		}
		response(w, 200, models.GlobalModel{Pais: x.Pais}, "user", err)
	} else {
		fmt.Println("Por favor vuelta a intentar cargar los datos en otro momento.")
		fmt.Println("")
	}
}
func ContinenteGet(w http.ResponseWriter, r *http.Request) {
	var x models.GlobalModel

	a := r.URL.Query()
	x.Continente.Nombre = a.Get("nombre")

	if x.Continente.Nombre != ""{
		globalDAO := globalDAO()
		x, err := globalDAO.GetOne(x, "continente")
		if err != nil {
			response(w, 404, x, "continente", err)
		}
		response(w, 200, x, "continente", nil)
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
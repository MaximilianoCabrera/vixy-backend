package actions

import (
	"net/http"

	"../models"
	"fmt"
	"strconv"
	"log"
)

func PaisCreate(w http.ResponseWriter, r *http.Request) {
	//var err error
	var x models.GlobalModel

	a := r.URL.Query()
	x.Pais.Nombre = a.Get("nombre")
	x.Continente.Nombre = a.Get("continente")
	x.Pais.Moneda = a.Get("moneda")
	x.Pais.Usohorario = a.Get("usoHorario")
	x.Pais.Idioma = a.Get("idioma")

	if x.Pais.Nombre != "" &&
		x.Continente.Nombre != "" &&
		x.Pais.Moneda != "" &&
		x.Pais.Usohorario != "" &&
		x.Pais.Idioma != "" {

		globalDAO := globalDAO()
		a, err := globalDAO.GetOne(x, "continente")
		checkErr("continente", "getOne", err)

		x.Pais.Continente = a.Continente[0].ID
		x, err = globalDAO.Create(&x, "pais")
		if err != nil {
			fmt.Println("No se pudo crear el pais.", err)
			response(w, 404, x, "pais", err)
		}
		response(w, 200, x, "pais", err)
	} else {
		fmt.Println("Por favor vuelta a intentar cargar los datos en otro momento.")
		fmt.Println("")
	}
}
func PaisGet(w http.ResponseWriter, r *http.Request) {
	x := models.GlobalModel{}

	a := r.URL.Query()
	x.Pais.Nombre = a.Get("nombre")
	x.Continente.Nombre = a.Get("continente")
	x.Pais.Moneda = a.Get("moneda")
	x.Pais.Usohorario = a.Get("usoHorario")
	x.Pais.Idioma = a.Get("idioma")

	if x.Pais.Nombre != "" ||
		x.Continente.Nombre != "" ||
		x.Pais.Moneda != "" ||
		x.Pais.Usohorario != "" ||
		x.Pais.Idioma != "" {

		globalDAO := globalDAO()
		if x.Continente.Nombre != ""{
			a, err := globalDAO.GetOne(x, "continente")
			checkErr("continente", "getOne", err)

			log.Println("continente: ", a.Continente[0].ID)
			x.Pais.Continente = a.Continente[0].ID
		}

		a, err := globalDAO.GetOne(x, "pais")
		if err != nil {
			responses(w, 404, a, "pais", err)
		}
		responses(w, 200, a, "pais", nil)
	} else {
		var x models.GlobalModels

		globalDAO := globalDAO()
		x, err := globalDAO.GetAll("pais")
		if err != nil {
			responses(w, 404, x, "pais", err)
		}
		responses(w, 200, x, "pais", nil)
	}
}
func PaisGetByID(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id, err := strconv.Atoi(params.Get("id"))
	msjError(err)

	globalDAO := globalDAO()
	x, err := globalDAO.GetByID(id, "pais")
	if err != nil {
		response(w, 404, x, "pais", err)
	}

	response(w, 200, x, "pais", nil)
}
func PaisUpdate(w http.ResponseWriter, r *http.Request) {
	var x models.GlobalModel
	a := r.URL.Query()

	if a.Get("id") == "" {
		fmt.Println("No ingresó ningún ID.")
	}

	id, err := strconv.Atoi(a.Get("id"))
	x.Pais.ID = id
	x.Pais.Nombre = a.Get("nombre")
	x.Continente.Nombre = a.Get("continente")
	x.Pais.Moneda = a.Get("moneda")
	x.Pais.Usohorario = a.Get("usoHorario")
	x.Pais.Idioma = a.Get("idioma")

	globalDAO := globalDAO()
	if x.Continente.Nombre != ""{
		a, err := globalDAO.GetOne(x, "continente")
		checkErr("continente", "getOne", err)

		x.Pais.Continente = a.Continente[0].ID
	}

	x, err = globalDAO.Update(x, "pais")
	if err != nil {
		response(w, 404, x, "pais", err)
	} else {
		response(w, 200, x, "pais", err)
	}
}
func PaisDelete(w http.ResponseWriter, r *http.Request) {
	var x models.GlobalModel
	params := r.URL.Query()
	id, err := strconv.Atoi(params.Get("id"))
	msjError(err)

	x.Pais.ID = id

	globalDAO := globalDAO()

	msj, err := globalDAO.Delete(&x, "pais")
	checkErr("pais", "eliminar", err)
	fmt.Print("ELiminado: ", msj)
}

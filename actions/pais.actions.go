package actions

import (
	"net/http"

	"../models"
	"fmt"
	"strconv"
)

//Controlar//
func PaisCreate(w http.ResponseWriter, r *http.Request) {
	var err error
	var pais models.Pais

	a := r.URL.Query()
	pais.Nombre = a.Get("nombre")
	pais.Continente = a.Get("continente")
	pais.Moneda = a.Get("moneda")
	pais.UsoHorario = a.Get("usoHorario")
	pais.Idioma = a.Get("idioma")

	if pais.Nombre != "" &&
		pais.Continente != "" &&
		pais.Moneda != "" &&
		pais.UsoHorario != "" &&
		pais.Idioma != "" {

		cont := models.Continente{Nombre: pais.Continente}
		globalDAO := globalDAO()
		continente := models.GlobalModel{Continente: cont}
		continente, err = globalDAO.GetOne(continente, "continente")
		checkErr("No se pudo obtener el continente.", err)

		x := models.GlobalModel{}
		x.Pais = pais
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
func PaisGet(w http.ResponseWriter, r *http.Request) {
	var pais models.Pais

	a := r.URL.Query()
	pais.Nombre = a.Get("nombre")
	pais.Continente = a.Get("continente")
	pais.Moneda = a.Get("moneda")
	pais.UsoHorario = a.Get("usoHorario")
	pais.Idioma = a.Get("idioma")

	if pais.Nombre != "" ||
		pais.Continente != "" ||
		pais.Moneda != "" ||
		pais.UsoHorario != "" ||
		pais.Idioma != "" {

		var x models.GlobalModel

		x.Pais = pais

		globalDAO := globalDAO()
		x, err := globalDAO.GetOne(x, "pais")
		if err != nil {
			response(w, 404, x, "pais", err)
		}
		response(w, 200, x, "pais", nil)
	} else {
		var x models.GlobalModels

		globalDAO := globalDAO()
		x, err := globalDAO.GetAll("pais")
		if err != nil {
			responses(w, 404, x, "pais", err)
		}
		responses(w, 200, x, "user", nil)
	}
}

func PaisGetByID(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id, err := strconv.Atoi(params.Get("id"))
	checkErr("Error: ", err)

	globalDAO := globalDAO()
	x, err := globalDAO.GetByID(id, "user")
	if err != nil {
		responseUser(w, 404, x, "user", err)
	}
	responseUser(w, 200, x, "user", nil)
}
func PaisGetOne(w http.ResponseWriter, r *http.Request) {
	var x models.GlobalModel
	a := r.URL.Query()
	x.User.Nombre = a.Get("nombre")
	x.User.Apellido = a.Get("apellido")
	x.User.Nick = a.Get("nick")
	x.User.Email = a.Get("email")

	globalDAO := globalDAO()
	x, err := globalDAO.GetOne(x, "user")
	if err != nil {
		responseUser(w, 404, x, "user", err)
	}
	responseUser(w, 200, x, "user", nil)
}
func PaisUpdate(w http.ResponseWriter, r *http.Request) {
	var x models.GlobalModel
	a := r.URL.Query()

	if a.Get("id") == "" {
		fmt.Println("No ingresó ningún ID.")
	}

	globalDAO := globalDAO()

	id, err := strconv.Atoi(a.Get("id"))
	x.User.ID = id
	x.User.Nombre = a.Get("nombre")
	x.User.Apellido = a.Get("apellido")
	x.User.Nick = a.Get("nick")
	x.User.Email = a.Get("email")
	x.User.Password = a.Get("password")

	img := models.Imagen{}
	if a.Get("imagen") != "" {
		// Cargo todos los valores de la imagen
		fmt.Println("Por obtener")
		x.Imagen.Imagen = a.Get("imagen")

		fmt.Println("Img: ", img)

		x, err = globalDAO.GetOne(x, "imagen")
		if err != nil {
			fmt.Print("Error: ", err)
		}
		x, err = globalDAO.Update(x, "imagen")
		if err != nil {
			fmt.Print("Error: ", err)
		}
	}

	x, err = globalDAO.Update(x, "user")
	if err != nil {
		responseUser(w, 404, x, "user", err)
	} else {
		responseUser(w, 200, x, "user", err)
	}
}
func PaisDelete(w http.ResponseWriter, r *http.Request) {
	var x models.GlobalModel
	params := r.URL.Query()
	id, err := strconv.Atoi(params.Get("id"))
	checkErr("Error: ", err)

	x.User.ID = id

	globalDAO := globalDAO()

	msj, err := globalDAO.Delete(&x, "user")
	checkErr("Error: ", err)
	fmt.Print("ELiminado: ", msj)
}

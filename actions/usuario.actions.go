package actions

import (
	"net/http"

	"../models"
	"fmt"
	"strconv"
	"log"
)

//Ok//
func UserCreate(w http.ResponseWriter, r *http.Request) {
	var trans1 = false
	var trans2 = false
	var err error
	var user models.Usuario

	a := r.URL.Query()
	user.Nombre = a.Get("nombre")
	user.Apellido = a.Get("apellido")
	user.Nick = a.Get("nick")
	user.Email = a.Get("email")
	user.Password = a.Get("password")
	user.TipoUsuario = 2

	if user.Nombre != "" &&
		user.Apellido != "" &&
		user.Nick != "" &&
		user.Email != "" &&
		user.Password != "" {

		var img models.Imagen
		if a.Get("imagen") != "" {
			// Cargo todos los valores de la imagen
			img.Imagen = a.Get("imagen")
		} else {
			img.Imagen = "user.jpg"
		}

		globalDAO := globalDAO()
		imagen := models.GlobalModel{}
		imagen.Imagen = img
		imagen, err = globalDAO.Create(&imagen, "imagen")
		checkErr("imagen", "crear", err)
		user.Imagen = imagen.Imagen.ID
		trans1 = true

		us := models.GlobalModel{}
		us.User = user
		us, err = globalDAO.Create(&us, "user")
		if err != nil {
			fmt.Println("No se pudo crear el usuario.", err)
			_, err = globalDAO.Delete(&imagen, "imagen")
			checkErr("imagen", "eliminar", err)
		} else {
			trans2 = true
		}

		if trans1 && trans2 == true {
			response(w, 200, models.GlobalModel{User: user}, "user", err)
		} else {
			response(w, 404, models.GlobalModel{User: user}, "user", err)
		}
	} else {
		fmt.Println("Por favor vuelta a intentar cargar los datos en otro momento.")
		fmt.Println("")
	}
}
func UserGet(w http.ResponseWriter, r *http.Request) {
	var user models.Usuario

	a := r.URL.Query()
	user.Nombre = a.Get("nombre")
	user.Apellido = a.Get("apellido")
	user.Nick = a.Get("nick")
	user.Email = a.Get("email")
	user.Password = a.Get("password")

	if user.Nombre != "" ||
		user.Apellido != "" ||
		user.Nick != "" ||
		user.Email != "" ||
		user.Password != "" {

		var x models.GlobalModel

		x.User = user
		globalDAO := globalDAO()
		fmt.Println("GETONE USER")
		a, err := globalDAO.GetBy(x, "user")
		if err != nil {
			fmt.Println("ERROR EN GETONE USER")
			responses(w, 404, a, "user", err)
		}

		for _, users := range a.User{
			log.Println(users)
		}
		responses(w, 200, a, "user", nil)
	} else {
		var x models.GlobalModels

		globalDAO := globalDAO()
		x, err := globalDAO.GetAll("user")
		if err != nil {
			responses(w, 404, x, "user", err)
		}
		responses(w, 200, x, "user", nil)
	}
}
func UserGetByID(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id, err := strconv.Atoi(params.Get("id"))
	msjError(err)

	globalDAO := globalDAO()
	x, err := globalDAO.GetByID(id, "user")
	if err != nil {
		response(w, 404, x, "user", err)
	}
	response(w, 200, x, "user", nil)
}
func UserUpdate(w http.ResponseWriter, r *http.Request) {
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

		a, err := globalDAO.GetBy(x, "imagen")
		if err != nil {
			responses(w, 404, a, "imagen",err)
		} else{
			responses(w, 200, a, "imagen", err)
		}

		x.Imagen = a.Imagen[0]

		x, err = globalDAO.Update(x, "imagen")
		checkErr("imagen", "update", err)
	}

	x, err = globalDAO.Update(x, "user")
	if err != nil {
		response(w, 404, x, "user", err)
	} else {
		response(w, 200, x, "user", err)
	}
}
func UserDelete(w http.ResponseWriter, r *http.Request) {
	var x models.GlobalModel
	params := r.URL.Query()
	id, err := strconv.Atoi(params.Get("id"))
	msjError(err)

	x.User.ID = id

	globalDAO := globalDAO()

	msj, err := globalDAO.Delete(&x, "user")
	checkErr("user", "eliminar", err)
	fmt.Print("ELiminado: ", msj)
}

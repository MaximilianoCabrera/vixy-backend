package actions

import (
	"net/http"

	"../dao/factory"
	"../dao/interfaces"
	"../models"
	"../utilities"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

// response Usuario
func responseUser(w http.ResponseWriter, status int, results models.GlobalModel, model string, err error) {
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
	}
}
// response Usuarios
func responseUsers(w http.ResponseWriter, status int, results models.GlobalModels, model string, err error) {
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
	}
}

// Func Errores
func msjError(err error) {
	fmt.Println("Error: ", err)
}
// Func config

func config()(models.Configuration){
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}
	return config
}

func globalDAO()(x interfaces.GlobalDAO){
	x = factory.GlobalFactoryDAO(config().Engine)
	return x
}
func checkErr(msj string, err error){
	log.Println(msj, err)
}

//Ok//
func UsuarioCreate(w http.ResponseWriter, r *http.Request) {

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
		checkErr("No se pudo cargar la imagen.", err)

		trans1 = true

		if trans1 == true {
			user.Imagen = imagen.Imagen.ID

			us := models.GlobalModel{}
			us.User = user
			us, err = globalDAO.Create(&us, "user")
			if err != nil {
				fmt.Println("No se pudo crear el usuario.", err)
				_, err = globalDAO.Delete(imagen.Imagen.ID, &imagen, "imagen")
				checkErr("Se han borrado tanto los datos del usuario como de la imagen.", err)
				trans2 = true
		}

		if trans1 && trans2 == true {
			responseUser(w, 200, models.GlobalModel{User:user}, "user", err)
		} else {
			fmt.Println("Por favor vuelta a intentar cargar los datos en otro momento.")
			fmt.Println("")
		}
	} else {
		responseUser(w, 404, models.GlobalModel{User: user}, "user", err)
		}
	}
}
//Revisar que se utilice el globalDAO()//
func UserGetAll(w http.ResponseWriter, r *http.Request) {
	var x models.GlobalModels
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Println(err)
	}
	globalDAO := factory.GlobalFactoryDAO(config.Engine)

	x, err = globalDAO.GetAll("user")
	if err != nil {
		responseUsers(w, 404, x, "user", err)
	}
	responseUsers(w, 200, x, "user", nil)
}
func UserGetByID(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		log.Println(err)
	}

	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Println(err)
	}
	globalDAO := factory.GlobalFactoryDAO(config.Engine)

	x, err := globalDAO.GetByID(id, "user")
	if err != nil {
		responseUser(w, 404, x, "user", err)
	}
	responseUser(w, 200, x, "user", nil)
}
func UserGetOne(w http.ResponseWriter, r *http.Request) {
	var x models.GlobalModel
	a := r.URL.Query()
	x.User.Nombre = a.Get("nombre")
	x.User.Apellido = a.Get("apellido")
	x.User.Nick = a.Get("nick")
	x.User.Email = a.Get("email")

	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}

	globalDAO := factory.GlobalFactoryDAO(config.Engine)
	x, err = globalDAO.GetOne(x, "user")
	if err != nil {
		responseUser(w, 404, x, "user", err)
	}
	responseUser(w, 200, x, "user", nil)
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

		x, err = globalDAO.GetOne(x, "imagen")
		if err != nil {
			fmt.Print("Error: ", err)
			//msjImagen = "No se pudo actualizar el usuario desde Actions."
		}
		x, err = globalDAO.Update(x, "imagen")
		if err != nil {
			fmt.Print("Error: ", err)
			//msjImagen = "No se pudo actualizar la imagen."
		}
	}

	x, err = globalDAO.Update(x, "user")
	if err != nil{
		//fmt.Println("Error: ", err)
		responseUser(w, 404, x, "user", err)
	} else{
		//fmt.Println("USUARIO: ", x)

		responseUser(w, 200, x, "user", err)
	}

}

//revisar//
//func UserDelete(){}
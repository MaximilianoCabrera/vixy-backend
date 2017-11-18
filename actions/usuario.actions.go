package actions

import (
	"net/http"

	"../dao/factory"
	"../models"
	"../utilities"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

// response Usuario
func responseUser(w http.ResponseWriter, status int, results models.Usuario, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(results)
}

// response Usuarios
func responseUsers(w http.ResponseWriter, status int, results []models.GlobalModel, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(results)
}
func msjError(err error) {
	fmt.Println("Error: ", err)
}

//Ok//
func UsuarioCreate(w http.ResponseWriter, r *http.Request) {
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}
	var trans1 = false
	var trans2 = false
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
		var msjUser = ""
		var msjImagen = ""

		var img models.Imagen
		if a.Get("imagen") != "" {
			// Cargo todos los valores de la imagen
			img.Imagen = a.Get("imagen")
		} else {
			img.Imagen = "user.jpg"
		}

		globalDAO := factory.GlobalFactoryDAO(config.Engine)
		imagen := models.GlobalModel{}
		imagen.Imagen = img
		imagen, err = globalDAO.Create(&imagen, "imagen")

		if err != nil {
			fmt.Print("Error: ", err)
			msjImagen = "No se pudo cargar la imagen."
		} else {
			msjImagen = "Imagen cargada correctamente"
			trans1 = true
		}

		if trans1 == true {
			user.Imagen = imagen.Imagen.ID

			us := models.GlobalModel{}
			us.User = user
			us, err = globalDAO.Create(&us, "user")
			if err != nil {
				fmt.Println("No se pudo crear el usuario.", err)
				msjImagen, err = globalDAO.Delete(imagen.Imagen.ID, &imagen, "imagen")
				if err != nil {
					fmt.Println("Error: ", err)
				}
				fmt.Println("Se han borrado tanto los datos del usuario como de la imagen.")
			}else{
				msjUser = "Usuario creado correctamente"
				trans2 = true
			}
		}

		if trans1 && trans2 == true {
			fmt.Println(msjImagen)
			fmt.Println(msjUser)
		} else {
			fmt.Println("Por favor vuelta a intentar cargar los datos en otro momento.")
			fmt.Println("")
		}
	} else {
		fmt.Println("Debe completar todos los campos")
	}
}
//Revisar//
func UserGetAll(w http.ResponseWriter, r *http.Request) {
	var x []models.GlobalModel
	//var users []models.Usuario
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}
	//userDAO := factory.UsuarioFactoryDAO(config.Engine)
	globalDAO := factory.GlobalFactoryDAO(config.Engine)

	x, err = globalDAO.GetAll("user")
	//users, err = userDAO.GetAll()
	fmt.Println(x)
	fmt.Println(err)
	if err != nil {
		responseUsers(w, 404, nil, err)
	}
	fmt.Println()
	responseUsers(w, 200, x, nil)
}
func UserGetOne(w http.ResponseWriter, r *http.Request) {

	var user models.Usuario
	a := r.URL.Query()
	user.Nombre = a.Get("nombre")
	user.Apellido = a.Get("apellido")
	user.Nick = a.Get("nick")
	user.Email = a.Get("email")

	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}

	if a.Get("tipoUsuario") != "" {
		tipoUsuarioDAO := factory.TipoUsuarioFactoryDAO(config.Engine)
		tipoUsuario := models.TipoUsuario{}
		tipoUsuario.Nombre = a.Get("tipoUsuario")

		tipoUser, err := tipoUsuarioDAO.GetOne(tipoUsuario)
		if err != nil {
			msjError(err)
		}
		user.TipoUsuario = tipoUser.ID
	}

	userDAO := factory.UsuarioFactoryDAO(config.Engine)
	user, err = userDAO.GetOne(user)
	if err != nil {
		responseUser(w, 404, user, err)
	}
	responseUser(w, 200, user, nil)
}
func UserGetBy(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()
	id, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		log.Fatalln(err)
	}

	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}
	userDAO := factory.UsuarioFactoryDAO(config.Engine)

	user, err := userDAO.GetByID(id)
	if err != nil {
		responseUser(w, 404, user, err)
	}
	responseUser(w, 200, user, nil)
}
func UserUpdate(w http.ResponseWriter, r *http.Request) {

	a := r.URL.Query()

	if a.Get("id") == "" {
		fmt.Println("No ingresó ningún ID.")
	}

	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}

	var user models.Usuario

	user.Nombre = a.Get("nombre")
	user.Apellido = a.Get("apellido")
	user.Nick = a.Get("nick")
	user.Email = a.Get("email")
	user.Password = a.Get("password")

	var msjUser= "No se pudo actualizar el usuario."
	var msjImagen= ""

	img := models.Imagen{}
	if a.Get("imagen") != "" {
		// Cargo todos los valores de la imagen
		fmt.Println("Por obtener")
		img.Imagen = a.Get("imagen")
		imagenDAO := factory.ImagenFactoryDAO(config.Engine)

		fmt.Println("Img: ", img)

		imagen, err := imagenDAO.GetOne(img)
		_, err = imagenDAO.Update(imagen)
		if err != nil {
			fmt.Print("Error: ", err)
			msjImagen = "No se pudo cargar la imagen."
		} else {
			fmt.Println("Imagen: ", imagen)
			user.ID = imagen.ID
			msjImagen = "Imagen cargada correctamente"
		}
	} else{
		fmt.Println("Imagen vacía")
	}

	// Cargo todos los valores del nuevo usuario

	usuarioDAO := factory.UsuarioFactoryDAO(config.Engine)
	userUpdated, err := usuarioDAO.Update(user)
	if err != nil {
		fmt.Println("")
		fmt.Println("No se pudo crear el usuario.", err)
		fmt.Println("")
	}

	fmt.Println("Usuario actualizado: ", userUpdated)
	fmt.Println("")
	fmt.Println(msjImagen)
	fmt.Println("")
	fmt.Println(msjUser)
}

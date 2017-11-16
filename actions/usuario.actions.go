package actions

import (
	"net/http"

	"../dao/factory"
	"../models"
	"../utilities"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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
func responseUsers(w http.ResponseWriter, status int, results []models.Usuario, err error) {
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

	usuarioDAO := factory.UsuarioFactoryDAO(config.Engine)
	usuario := models.Usuario{}
	imagenDAO := factory.ImagenFactoryDAO(config.Engine)
	img := models.Imagen{}

	fmt.Print("Nombre: ")
	fmt.Scan(&usuario.Nombre)
	fmt.Print("Apellido: ")
	fmt.Scan(&usuario.Apellido)
	fmt.Print("Nick: ")
	fmt.Scan(&usuario.Nick)
	fmt.Print("Email: ")
	fmt.Scan(&usuario.Email)
	fmt.Print("Password: ")
	fmt.Scan(&usuario.Pass)
	usuario.TipoUsuario = 2
	fmt.Print("Nombre de la imagen: ")
	fmt.Scan(&img.Imagen)

	// Cargo todos los valores de la imagen

	var msjUser = "No se pudo cargar el usuario."
	var msjImagen = ""

	imagen, err := imagenDAO.Create(&img)
	if err != nil {
		fmt.Print("Error: ", err)
		msjImagen = "No se pudo cargar la imagen."
	} else {
		msjImagen = "Imagen cargada correctamente"
		trans1 = true
	}

	// Cargo todos los valores del nuevo usuario
	if trans1 == true {
		usuario.IDImagen = imagen.ID
		fmt.Println("*******")
		fmt.Println(imagen.ID)
		fmt.Println("*******")
		msjUser, err = usuarioDAO.Create(&usuario)
		if err != nil {
			fmt.Println("")
			fmt.Println("No se pudo crear el usuario.", err)
			fmt.Println("")

			err2 := imagenDAO.Delete(imagen.ID)
			fmt.Println("Se han borrado tanto los datos del usuario como de la imagen.")
			fmt.Println("")
			if err2 != nil {
				fmt.Println("Error2: ", err2)
				fmt.Println("")
			}
		}
		fmt.Println("*******")
		fmt.Println(usuario)
		fmt.Println("*******")

		trans2 = true
	}

	if trans1 && trans2 == true {
		fmt.Println(msjImagen, msjUser)
	} else {
		fmt.Println("Por favor vuelta a intentar cargar los datos en otro momento.")
		fmt.Println("")
	}
}
func UserGetAll(w http.ResponseWriter, r *http.Request) {
	var users []models.Usuario
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}
	userDAO := factory.UsuarioFactoryDAO(config.Engine)

	users, err = userDAO.GetAll()
	fmt.Println(users)
	fmt.Println(err)
	if err != nil {
		responseUsers(w, 404, nil, err)
	}
	responseUsers(w, 200, users, nil)
}

//Revisar//

func UserGetOne(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.URL.Path)


/*
	params := mux.Vars(r)
	user := models.Usuario{}

	//Obtengo ID - No creo q lo use.
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		msjError(err)
	}
	user.ID = id

	user.Nombre = params["nombre"]
	user.Apellido = params["apellido"]
	user.Email = params["mail"]

	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}

	tipoUsuarioDAO := factory.TipoUsuarioFactoryDAO(config.Engine)
	tipoUsuario := models.TipoUsuario{}
	tipoUsuario.Nombre = params["Usuario"]
	tipoUser, err := tipoUsuarioDAO.GetOne(tipoUsuario)
	if err != nil {
		msjError(err)
	}
	user.TipoUsuario = tipoUser.ID

	img, err := strconv.Atoi(params["imagen"])
	if err != nil {
		msjError(err)
	}
	user.IDImagen = img

	userDAO := factory.UsuarioFactoryDAO(config.Engine)
	user, err = userDAO.GetOne(user)
	if err != nil {
		msjError(err)
		//responseUser(w, 404, user, err)
	}
	responseUser(w, 200, user, nil)
*/
}
func UserGetBy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	params := models.Usuario{}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Error: ", err)
	}
	params.ID = id

	if err != nil {
		log.Fatalln(err)
	}

	var user models.Usuario

	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}
	userDAO := factory.UsuarioFactoryDAO(config.Engine)

	user, err = userDAO.GetByID(id)
	if err != nil {
		fmt.Println("ERROR")
		//responseUser(w, 404, user, err)
	}
	responseUser(w, 200, user, nil)
}

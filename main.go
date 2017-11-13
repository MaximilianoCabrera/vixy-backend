package main

import (
	"log"
	"net/http"
	"./routes"
	"github.com/rs/cors"
)


func main() {
	/*config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}*/
		router := routes.NewMainRouter()

		handler := cors.Default().Handler(router)
		server := http.ListenAndServe(":8080", handler)
		log.Fatal(server)


		//usa la impletentación específica
		//userDAO := factory.FactoryDAO(config.Engine)
/*
	var user = models.User{}

	//MUESTRO LOS TODOS USUARIOS
	fmt.Println("")
	fmt.Println("GetAll()")
	fmt.Println(userDAO.GetAll())
	fmt.Println("")
	fmt.Println("**********************************")

	//ELIJO Y MUESTRO UN USUARIO
	fmt.Println("")
	fmt.Println("GetById()")
	fmt.Print("Ingrese el id del usuario: ")
	fmt.Scan(&user.ID)
	fmt.Println(userDAO.GetByID(user.ID))
	fmt.Println("")
	fmt.Println("**********************************")

	//ACTUALIZO EL USUARIO SELECCIONADO
	fmt.Println("")
	fmt.Println("Update()")
	fmt.Println("")
	fmt.Print("Id del Usuario a Modificar: ")
	fmt.Print(user.ID)

	fmt.Println()

	fmt.Print(user.Nombre)
	fmt.Print("Nuevo Nombre: ")
	fmt.Scan(&user.Nombre)
	fmt.Print("Apellido: ")
	fmt.Scan(&user.Apellido)
	fmt.Print("Correo: ")
	fmt.Scan(&user.Email)

	fmt.Println(userDAO.Update(user))

	fmt.Println("")
	fmt.Println("")
	fmt.Println("Nuevos Datos: ")
	fmt.Println(userDAO.GetByID(user.ID))
	fmt.Println("")
	fmt.Println("**********************************")

	//ELIMINO EL USUARIOS SELECCIONADO
	fmt.Println("")
	fmt.Println("Delete()")
	fmt.Println("")
	fmt.Println(userDAO.GetAll())
	fmt.Println("")
	fmt.Print("Ingrese el ID del usuario que desea elminar: ")
	fmt.Scan(&user.ID)

	fmt.Println(userDAO.Delete(user.ID))

	fmt.Println("")
	fmt.Printf("El usuario con ID %s fue eliminado correctamente", user.ID)
	fmt.Println("")
	fmt.Println("**********************************")

	//CREO UN USUARIO
	fmt.Println("")
	fmt.Println("Create()")

	user = models.User{
		Nombre:   "Alfredo",
		Apellido: "Lopez",
		Nick: "Alfred",
		Email:    "AlfredoLopez@gmail.com",
		Password: "contraseña",
	}

	fmt.Println(userDAO.Create(user))
	fmt.Println("Usuario creado correctamente")
	fmt.Println(userDAO.GetAll())
*/
}

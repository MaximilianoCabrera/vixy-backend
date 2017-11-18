package mysql

import (
	"../../models"
	"log"
	"fmt"
)

type GlobalImplMysql struct{}

//TODO: Agregar los DB().BEGIN() y COMMIT
func checkErr(msj string, err error) {
	if err != nil {
		log.Println(msj, err)
	}
}

//ok//
func (dao GlobalImplMysql) Create(x *models.GlobalModel, model string) (models.GlobalModel, error) {
	switch model {
	case "user":
		err := DB().CreateAndRead(&x.User)
		checkErr("Error al crear Usuario: ", err)
	case "imagen":
		err := DB().CreateAndRead(&x.Imagen)
		checkErr("Error al crear imagen: ", err)
	default:
		log.Println("Modelo ingresado no existente")
	}
	return *x, nil
}
func (dao GlobalImplMysql) Delete(id int, x *models.GlobalModel, model string) (string, error) {
	var msj = ""
	var err error

	switch model {
	case "user":
		err = DB().Delete(x.User)
		checkErr("Error al eliminar Usuario: ", err)
		msj = "Usuario eliminado correctamente."
	case "imagen":
		err = DB().Delete(x.Imagen)
		checkErr("Error al eliminar Imagen: ", err)
		msj = "Imagen eliminado correctamente."
	default:
		msj = "Modelo ingresado no existente"
	}

	return msj, err
}
func (dao GlobalImplMysql) GetAll(model string) (models.GlobalModels, error){
	var x models.GlobalModel
	var a models.GlobalModels

	switch model {
	case "user":
		err := DB().Read(&a.User, "SELECT * FROM usuario")
		checkErr("Error al querer obtener todas los usuarios", err)

		for _, usuario := range a.User{
			x.User = usuario
		}

	case "imagen":
		err := DB().Read(&a.Imagen, "SELECT * FROM imagen")
		checkErr("Error al querer obtener todas las imágenes: ", err)

		for _, img := range a.Imagen{
			x.Imagen = img
		}
	default:
		log.Println("Modelo ingresado no existente")
	}

	return a, nil
}
func (dao GlobalImplMysql) GetByID(id int, model string) (models.GlobalModel, error){
	var x models.GlobalModel
	//var err error

	switch model {
	case "user":
		var user models.Usuario
		err := DB().Read(&user,"SELECT * FROM usuario WHERE id = ?", id)
		checkErr("Error al obtener usuario: ", err)

		x.User = user
	case "imagen":
		var imagen models.Imagen
		err := DB().Read(&imagen, "SELECT * FROM imagen WHERE id = ?", id)
		checkErr("Error al obtener imagen: ", err)
	default:
		log.Println("Modelo ingresado no existente")
	}
	return x, nil
}
func (dao GlobalImplMysql) GetOne (x models.GlobalModel, model string) (models.GlobalModel, error){
	var err error
	switch model {
	case "user":
		var user models.Usuario
		err := DB().Read(&user, "SELECT * FROM usuario WHERE nombre = ? OR apellido = ? OR nick = ? OR email = ? ", x.User.Nombre, x.User.Apellido, x.User.Nick, x.User.Email)

		x.User = user
		checkErr("Error al obtener Usuario: ", err)
	case "imagen":
		checkErr("Error al crear imagen: ", err)
	default:
		log.Println("Modelo ingresado no existente")
	}
	return x, nil
}

//revisar//
func (dao GlobalImplMysql) Update (x models.GlobalModel, model string) (models.GlobalModel, error){
	var err error
	switch model {
	case "user":
		var user models.Usuario
		//fmt.Println("SELECT * FROM usuario WHERE id = ?", x.User.ID)
		err := DB().Read(&user, "SELECT * FROM usuario WHERE id = ?", x.User.ID)
		checkErr("Error al buscar el Usuario a modificar: ", err)

		if x.User.Nombre == ""{
			fmt.Println("Nombre: ", x.User.Nombre)
			x.User.Nombre = user.Nombre
		}
		if x.User.Apellido == ""{
			x.User.Apellido = user.Apellido
		}
		if x.User.Nick == ""{
			x.User.Nick = user.Nick
		}
		if x.User.Email == ""{
			x.User.Email = user.Email
		}
		if x.User.Password == ""{
			x.User.Password = user.Password
		}
		x.User.TipoUsuario = user.TipoUsuario
		x.User.Imagen = user.Imagen

		err = DB().Update(x.User)
		checkErr("No se pudo actualizar el usuario desde IMPL: ", err)

	case "imagen":

		checkErr("Error al crear imagen: ", err)
	default:
		log.Println("Modelo ingresado no existente")
	}
	return x, nil
}


package mysql

import (
	"../../models"
	"log"
	"fmt"
)

type GlobalImplMysql struct{}
var msjreturn = ""
//TODO: Agregar los DB().BEGIN() y COMMIT
func checkErr(msj string, err error) {
	if err != nil {
		log.Println(msj, err)
		msjreturn = "No se encontró el archivo indicado."
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
	case "pais":
		err := DB().CreateAndRead(&x.Pais)
		checkErr("Error al crear pais: ", err)
	default:
		log.Println("Modelo ingresado no existente")
	}
	return *x, nil
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
	case "pais":
		err := DB().Read(&a.Pais, "SELECT * FROM pais")
		checkErr("Error al querer obtener todas los paises", err)

		for _, pais := range a.Pais{
			x.Pais = pais
		}
	default:
		log.Println("Modelo ingresado no existente")
	}

	return a, nil
}
func (dao GlobalImplMysql) GetByID(id int, model string) (models.GlobalModel, error){
	var x models.GlobalModel

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
		x.Imagen = imagen
	case "pais":
		var pais models.Pais
		err := DB().Read(&pais, "SELECT * FROM pais WHERE id = ?")
		checkErr("Error al obtener pais: ", err)
		x.Pais = pais
	default:
		log.Println("Modelo ingresado no existente")
	}
	return x, nil
}
func (dao GlobalImplMysql) GetOne(x models.GlobalModel, model string) (models.GlobalModel, error){
	switch model {
	case "user":
		var user models.Usuario
		err := DB().Read(&user, "SELECT * FROM usuario WHERE nombre = ? OR apellido = ? OR nick = ? OR email = ? ", x.User.Nombre, x.User.Apellido, x.User.Nick, x.User.Email)
		x.User = user
		checkErr("Error al obtener Usuario: ", err)
	case "imagen":
		var img models.Imagen
		err := DB().Read(&img, "SELECT * FROM pais WHERE imagen = ?", x.Imagen.Imagen)
		checkErr("Error al crear imagen: ", err)
		x.Imagen = img
	case "pais":
		var pais models.Pais
		err := DB().Read(&pais, "SELECT * FROM pais WHERE nombre = ?, idContinente = ?, moneda = ?, usoHorario = ?, idioma = ?", x.Pais.Nombre, x.Pais.Continente, x.Pais.Moneda, x.Pais.UsoHorario, x.Pais.Idioma)
		checkErr("Error al crear pais: ", err)
		x.Pais = pais
	default:
		log.Println("Modelo ingresado no existente")
	}
	return x, nil
}
func (dao GlobalImplMysql) Update (x models.GlobalModel, model string) (models.GlobalModel, error){
	switch model {
	case "user":
		var user models.Usuario

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
		var img models.Imagen

		err := DB().Read(&img, "SELECT * FROM imagen WHERE id = ?", x.Imagen.ID)
		checkErr("Error al buscar la Imagen a modificar: ", err)

		if x.Imagen.Imagen == ""{
			x.Imagen.Imagen = img.Imagen
		}
		err = DB().Update(x.Imagen)

		checkErr("No se pudo actualizar la imagen desde IMPL: ", err)
	case "pais":
		var pais models.Pais

		err := DB().Read(&pais, "SELECT * FROM pais WHERE id = ?", x.Pais.ID)
		checkErr("Error al buscar el Pais a modificar: ", err)

		if x.Pais.Nombre == ""{
			x.Pais.Nombre = pais.Nombre
		}
		if x.Pais.Continente == ""{
			x.Pais.Continente = pais.Continente
		}
		if x.Pais.Moneda == ""{
			x.Pais.Moneda = pais.Moneda
		}
		if x.Pais.UsoHorario == ""{
			x.Pais.UsoHorario = pais.UsoHorario
		}
		if x.Pais.Idioma == ""{
			x.Pais.Idioma = pais.Idioma
		}

		err = DB().Update(x.Pais)
		checkErr("No se pudo actualizar el pais desde IMPL: ", err)
	default:
		log.Println("Modelo ingresado no existente")
	}
	return x, nil
}
func (dao GlobalImplMysql) Delete(x *models.GlobalModel, model string) (string, error) {
	var err error

	switch model {
	case "user":
		fmt.Println("Eliminando desde el IMPL")
		msjreturn = "Usuario eliminado correctamente."
		err = DB().Delete(x.User)
		checkErr("Error al eliminar Usuario: ", err)
	case "imagen":
		msjreturn = "Imagen eliminada correctamente."
		err = DB().Delete(x.Imagen)
		checkErr("Error al eliminar Imagen: ", err)
	case "pais":
		msjreturn = "Pais eliminado correctamente."
		err = DB().Delete(x.Pais)
		checkErr("Error al eliminar Pais: ", err)
	default:
		msjreturn = "Modelo ingresado no existente"
	}

	return msjreturn, err
}

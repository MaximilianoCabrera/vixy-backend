package mysql

import (
	"../../models"
	"fmt"
	"log"
)

type GlobalImplMysql struct{}

func checkErr(model string, accion string, err error) {
	if err != nil {
		log.Println("No se pudo "+ accion + " " + model + ". Error desde IMPL: ", err)
		/*
		switch model {
		case "actividad":
		case "ciudad":
		case "comentarioEntrada":
		case "continente":
			log.Println("No se pudo "+accion+" el Continente. Error desde IMPL: ", err)
		case "entrada":
		case "imagen":
			log.Println("No se pudo "+accion+" la Imagen. Error desde IMPL: ", err)
		case "imagenCiudad":
		case "imagenEntrada":
		case "imagenPais":
		case "pais":
			log.Println("No se pudo "+accion+" el Pais. Error desde IMPL: ", err)
		case "tipoUsuario":
		case "topic":
		case "user":
			log.Println("No se pudo "+accion+" el Usuario. Error desde IMPL: ", err)
		case "usuarioEntrada":

		default:
			log.Println("Error desde IMPL: ", err)
		}
		*/
	}
}

//TODO: Agregar los DB().BEGIN() y COMMIT
func (dao GlobalImplMysql) Create(x *models.GlobalModel, model string) (models.GlobalModel, error) {
	switch model {
	case "actividad":
		err := DB().CreateAndRead(&x.Actividad)
		checkErr("actividad", "crear", err)
	case "ciudad":
		err := DB().CreateAndRead(&x.Ciudad)
		checkErr("ciudad", "crear", err)
	case "comentario":
		err := DB().CreateAndRead(&x.Comentario)
		checkErr("comentario", "crear", err)
	case "comentarioEntrada":
		err := DB().CreateAndRead(&x.ComentarioEntrada)
		checkErr("comentarioEntrada", "crear", err)
	case "continente":
		err := DB().CreateAndRead(&x.Continente)
		checkErr("continente", "crear", err)
	case "entrada":
		err := DB().CreateAndRead(&x.Entrada)
		checkErr("entrada", "crear", err)
	case "imagen":
		err := DB().CreateAndRead(&x.Imagen)
		checkErr("imagen", "crear", err)
	case "imagenCiudad":
		err := DB().CreateAndRead(&x.ImagenCiudad)
		checkErr("imagenCiudad", "crear", err)
	case "imagenEntrada":
		err := DB().CreateAndRead(&x.ImagenEntrada)
		checkErr("imagenEntrada", "crear", err)
	case "imagenPais":
		err := DB().CreateAndRead(&x.ImagenPais)
		checkErr("imagenPais", "crear", err)
	case "pais":
		err := DB().CreateAndRead(&x.Pais)
		checkErr("pais", "crear", err)
	case "tipoUsuario":
		err := DB().CreateAndRead(&x.TipoUsuario)
		checkErr("tipoUsuario", "crear", err)
	case "topic":
		err := DB().CreateAndRead(&x.Topic)
		checkErr("topic", "crear", err)
	case "user":
		err := DB().CreateAndRead(&x.User)
		checkErr("user", "crear", err)
	case "usuarioEntrada":
		err := DB().CreateAndRead(&x.UsuarioEntrada)
		checkErr("usuarioEntrada", "crear", err)
	default:
		log.Println("Modelo ingresado no existente")
	}
	return *x, nil
}
func (dao GlobalImplMysql) GetAll(model string) (models.GlobalModels, error) {
	var a models.GlobalModels

	switch model {
	case "actividad":
		err := DB().Read(&a.Actividad, "SELECT * FROM actividad")
		checkErr("actividad", "getAll", err)
	case "ciudad":
		err := DB().Read(&a.Ciudad, "SELECT * FROM ciudad")
		checkErr("ciudad", "getAll", err)
	case "comentario":
		err := DB().Read(&a.Comentario, "SELECT * FROM comentario")
		checkErr("comentario", "getAll", err)
	case "comentarioEntrada":
		err := DB().Read(&a.ComentarioEntrada, "SELECT * FROM comentarioEntrada")
		checkErr("comentarioEntrada", "getAll", err)
	case "continente":
		err := DB().Read(&a.Continente, "SELECT * FROM continente")
		checkErr("continente", "getAll", err)
	case "entrada":
		err := DB().Read(&a.Entrada, "SELECT * FROM entrada")
		checkErr("entrada", "getAll", err)
	case "imagen":
		err := DB().Read(&a.Imagen, "SELECT * FROM imagen")
		checkErr("imagen", "getAll", err)
	case "imagenCiudad":
		err := DB().Read(&a.ImagenCiudad, "SELECT * FROM imagenCiudad")
		checkErr("imagenCiudad", "getAll", err)
	case "imagenEntrada":
		err := DB().Read(&a.ImagenEntrada, "SELECT * FROM imagenEntrada")
		checkErr("imagenEntrada", "getAll", err)
	case "imagenPais":
		err := DB().Read(&a.ImagenPais, "SELECT * FROM imagenPais")
		checkErr("imagenPais", "getAll", err)
	case "pais":
		err := DB().Read(&a.Pais, "SELECT * FROM pais")
		checkErr("pais", "getAll", err)
	case "tipoUsuario":
		err := DB().Read(&a.TipoUsuario, "SELECT * FROM tipoUsuario")
		checkErr("tipoUsuario", "getAll", err)
	case "topic":
		err := DB().Read(&a.Topic, "SELECT * FROM topic")
		checkErr("topic", "getAll", err)
	case "user":
		err := DB().Read(&a.User, "SELECT * FROM usuario")
		checkErr("user", "getAll", err)
	case "usuarioEntrada":
		err := DB().Read(&a.UsuarioEntrada, "SELECT * FROM usuarioEntrada")
		checkErr("usuarioEntrada", "getAll", err)
	default:
		log.Println("Modelo ingresado no existente")
	}
	return a, nil
}

func (dao GlobalImplMysql) GetByID(id int, model string) (models.GlobalModel, error) {
	var x models.GlobalModel

	switch model {
	case "actividad":
	case "ciudad":
	case "comentarioEntrada":
	case "continente":
		err := DB().Read(&x.Continente, "SELECT * FROM continente WHERE id = ?", id)
		checkErr("continente", "getByID", err)
	case "entrada":
	case "imagen":
		err := DB().Read(&x.Imagen, "SELECT * FROM imagen WHERE id = ?", id)
		checkErr("imagen", "getByID", err)
	case "imagenCiudad":
	case "imagenEntrada":
	case "imagenPais":
	case "pais":
		err := DB().Read(&x.Pais, "SELECT * FROM pais WHERE id = ?", id)
		checkErr("pais", "getByID", err)
	case "tipoUsuario":
	case "topic":
	case "user":
		err := DB().Read(&x.User, "SELECT * FROM usuario WHERE id = ?", id)
		checkErr("user", "getByID", err)
	case "usuarioEntrada":
	default:
		log.Println("Modelo ingresado no existente")
	}
	return x, nil
}
func (dao GlobalImplMysql) GetBy(x models.GlobalModel, model string) (models.GlobalModels, error) {
	a := models.GlobalModels{}
	switch model {
	case "actividad":
	case "ciudad":
	case "comentarioEntrada":
	case "continente":
		err := DB().Read(&a.Continente, "SELECT * FROM continente WHERE nombre = ?", x.Continente.Nombre)
		checkErr("continente", "getOne", err)
	case "entrada":
	case "imagen":
		err := DB().Read(&a.Imagen, "SELECT * FROM pais WHERE imagen = ?", x.Imagen.Imagen)
		checkErr("imagen", "getOne", err)
	case "imagenCiudad":
	case "imagenEntrada":
	case "imagenPais":
	case "pais":
		err := DB().Read(&a.Pais, "SELECT * FROM pais WHERE nombre = ? OR continente = ? OR moneda = ? OR usoHorario = ? OR idioma = ?", x.Pais.Nombre, x.Continente.ID, x.Pais.Moneda, x.Pais.Usohorario, x.Pais.Idioma)
		checkErr("pais", "getOne", err)
	case "tipoUsuario":
	case "topic":
	case "user":
		err := DB().Read(&a.User, "SELECT * FROM usuario WHERE nombre = ? OR apellido = ? OR nick = ? OR email = ? ", x.User.Nombre, x.User.Apellido, x.User.Nick, x.User.Email)
		checkErr("user", "getOne", err)
	case "usuarioEntrada":
	default:
		log.Println("Modelo ingresado no existente")
	}
	fmt.Println("MSJ: ", a)
	return a, nil
}
func (dao GlobalImplMysql) Update(x models.GlobalModel, model string) (models.GlobalModel, error) {
	switch model {
	case "actividad":
	case "ciudad":
	case "comentarioEntrada":
	case "continente":
	case "entrada":
	case "imagen":
		var img models.Imagen

		err := DB().Read(&img, "SELECT * FROM imagen WHERE id = ?", x.Imagen.ID)
		checkErr("imagen", "getByID", err)

		if x.Imagen.Imagen == "" {
			x.Imagen.Imagen = img.Imagen
		}
		err = DB().Update(x.Imagen)
		checkErr("imagen", "actualizar", err)
	case "imagenCiudad":
	case "imagenEntrada":
	case "imagenPais":
	case "pais":
		var pais models.Pais

		err := DB().Read(&pais, "SELECT * FROM pais WHERE id = ?", x.Pais.ID)
		checkErr("pais", "getByID", err)

		if x.Pais.Nombre == "" {
			x.Pais.Nombre = pais.Nombre
		}
		if x.Pais.Continente == 0 {
			x.Pais.Continente = pais.Continente
		}
		if x.Pais.Moneda == "" {
			x.Pais.Moneda = pais.Moneda
		}
		if x.Pais.Usohorario == "" {
			x.Pais.Usohorario = pais.Usohorario
		}
		if x.Pais.Idioma == "" {
			x.Pais.Idioma = pais.Idioma
		}

		err = DB().Update(x.Pais)
		checkErr("pais", "actualizar", err)
	case "tipoUsuario":
	case "topic":
	case "user":
		var user models.Usuario
		err := DB().Read(&user, "SELECT * FROM usuario WHERE id = ?", x.User.ID)
		checkErr("user", "getByID", err)

		if x.User.Nombre == "" {
			x.User.Nombre = user.Nombre
		}
		if x.User.Apellido == "" {
			x.User.Apellido = user.Apellido
		}
		if x.User.Nick == "" {
			x.User.Nick = user.Nick
		}
		if x.User.Email == "" {
			x.User.Email = user.Email
		}
		if x.User.Password == "" {
			x.User.Password = user.Password
		}
		x.User.TipoUsuario = user.TipoUsuario
		x.User.Imagen = user.Imagen

		err = DB().Update(x.User)
		checkErr("user", "actualizar", err)
	case "usuarioEntrada":
	default:
		log.Println("Modelo ingresado no existente")
	}
	return x, nil
}
func (dao GlobalImplMysql) Delete(x *models.GlobalModel, model string) (string, error) {
	var err error
	var msjReturn = ""
	switch model {
	case "actividad":
	case "ciudad":
	case "comentarioEntrada":
	case "continente":
	case "entrada":
	case "imagen":
		msjReturn = "Imagen eliminada correctamente."
		err = DB().Delete(x.Imagen)
		checkErr("imagen", "eliminar", err)
	case "imagenCiudad":
	case "imagenEntrada":
	case "imagenPais":
	case "pais":
		msjReturn = "Pais eliminado correctamente."
		err = DB().Delete(x.Pais)
		checkErr("pais", "eliminar", err)
	case "tipoUsuario":
	case "topic":
	case "user":
		msjReturn = "Usuario eliminado correctamente."
		err = DB().Delete(x.User)
		checkErr("user", "eliminar", err)
	case "usuarioEntrada":
	default:
		msjReturn = "Modelo ingresado no 	existente"
	}
	return msjReturn, err
}

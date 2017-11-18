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
		log.Fatal(msj, err)
		panic(err)
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

//revisar//
func (dao GlobalImplMysql) GetAll(model string) ([]models.GlobalModel, error){
	//TODO: Arreglar aca!!
	var err error
	fmt.Println("GETALL()")
	var a *models.GlobalModel

	switch model {
	case "user":
		fmt.Println("A: ", a)
		err := DB().Read(&a)
		checkErr("Error al buscar los Usuario: ", err)
		fmt.Println(a)

		//a.User = append(a.User, users)

	case "imagen":

		checkErr("Error al buscar las imagen: ", err)
	default:
		log.Println("Modelo ingresado no existente")
	}
	var x []models.GlobalModel
	x = append(x, *a)

	return x, nil
}
func (dao GlobalImplMysql) GetByID(id int, model string) (models.GlobalModel, error){
	var err error
	var x models.GlobalModel
	switch model {
	case "user":
		checkErr("Error al crear Usuario: ", err)
	case "imagen":

		checkErr("Error al crear imagen: ", err)
	default:
		log.Println("Modelo ingresado no existente")
	}
	return x, nil
}
func (dao GlobalImplMysql) GetOne (x models.GlobalModel, model string) (models.GlobalModel, error){
	var err error
	switch model {
	case "user":

		checkErr("Error al crear Usuario: ", err)
	case "imagen":

		checkErr("Error al crear imagen: ", err)
	default:
		log.Println("Modelo ingresado no existente")
	}
	return x, nil
}
func (dao GlobalImplMysql) Update (x models.GlobalModel, model string) (models.GlobalModel, error){
	var err error
	switch model {
	case "user":

		checkErr("Error al crear Usuario: ", err)
	case "imagen":

		checkErr("Error al crear imagen: ", err)
	default:
		log.Println("Modelo ingresado no existente")
	}
	return x, nil
}


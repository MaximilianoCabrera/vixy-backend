package mysql

import (
	//"log"
	//"errors"
	"../../models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

)

type UsuarioImplMysql struct{}
//TODO: Agregar los DB().BEGIN() y COMMIT
//OK//

func (dao UsuarioImplMysql) Create(u *models.Usuario) (string, error){
	err := DB().Create(u)
	checkErr("nose", err)

	var msj = "Se pudo cargar el usuario y la imagen correctamente."
	return msj, nil
}
//Probar//
func (dao UsuarioImplMysql) GetAll() ([]models.Usuario, error) {

	var usuarios []models.Usuario
	err := DB().Read(usuarios,"SELECT * FROM usuario")
	if err != nil{
		fmt.Println("ERROR 1: ", err)
	}

	fmt.Println(len(usuarios))


	//query := "SELECT * FROM usuario"
	//var usuarios []models.Usuario

/*
	//usuarios := make([]models.Usuario, 0)
	db := get()
	defer db.Close()

	//creo una sentencia=statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return usuarios, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return usuarios, err
	}

	for rows.Next() {
		var row models.Usuario
		err := rows.Scan(&row.ID, &row.Nombre, &row.Apellido, &row.Nick, &row.Email, &row.Pass, &row.TipoUsuario, &row.IDImagen)

		if err != nil {
			return usuarios, err
		}

		usuarios = append(usuarios, row)
	}
*/
	return usuarios, nil
}
func (dao UsuarioImplMysql) GetByID(id int) (models.Usuario, error) {

	user := &models.Usuario{}
	fmt.Println("ID: ", id)
	err := DB().Read(user, "SELECT * FROM usuario WHERE id = ?", 1)
	if err != nil{
		fmt.Println("error: ", err)
	}
	var u = *user
	fmt.Println(user.Nombre)

	return u, nil

}
func (dao UsuarioImplMysql) GetOne(u models.Usuario) (models.Usuario, error){
	var usuario = models.Usuario{}
/*
	query := "SELECT id, nombre, apellido, nick, email, password, idTipoUsuario, idImagen FROM usuario " +
		"WHERE id = ? " +
		"OR nombre = ? " +
		"OR apellido = ? " +
		"OR nick = ? " +
		"OR email = ? " +
		"OR idTipoUsuario = ? "

	db := get()
	defer db.Close()

	//creo una sentencia=statement
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	row := db.QueryRow(query, u.ID, u.Nombre, u.Apellido, u.Nick, u.Email, u.TipoUsuario)

	err = row.Scan(&usuario.ID, &usuario.Nombre, &usuario.Apellido, &usuario.Nick, &usuario.Email, &usuario.Password, &usuario.TipoUsuario, &usuario.Imagen)
	if err != nil {
		return usuario, err
	}
*/
	return usuario, nil
}
func (dao UsuarioImplMysql) Update(u models.Usuario) (models.Usuario, error) {

	fmt.Println("Actualizando Usuario")
	fmt.Println("...")
	user := &models.Usuario{}
	fmt.Println("Obtengo los datos del viejo user")
	fmt.Println("...")
	err := DB().Read(user, "SELECT * FROM usuarios WHERE id = ?", u.ID)
	fmt.Println("Viejo user: ", user)
	fmt.Println("...")


	fmt.Println("Actualizo usuario")
	fmt.Println("...")
	user.Nombre = u.Nombre

	fmt.Println("Nuevo Usuario", user)
	fmt.Println("...")
	err = DB().Update(user)

	fmt.Println("Usuario actualizado", user)
	fmt.Println("...")

	return u, err
/*
	fmt.Println("Actualizando Usuario")
	fmt.Println("...")

	db := get()
	defer db.Close()

	query := "UPDATE usuario SET nombre = ?, apellido = ?, nick = ?, email = ?, password = ?, idImagen = ? WHERE id = ?"

	stmt, err := db.Prepare(query)
	checkErr(err)

	defer stmt.Close()

	row, err := stmt.Exec(u.Nombre, u.Apellido, u.Nick, u.Email, u.Pass, u.IDImagen, u.ID)
	checkErr(err)

	affect, err := row.RowsAffected()
	checkErr(err)

	fmt.Println("Afected: ", affect)

	return u, nil
*/
}
func (dao UsuarioImplMysql) Delete(id int) error {
/*
	query := "DELETE FROM usuario WHERE id = ?"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	i, _ := row.RowsAffected()
	if i != 1 {
		return errors.New("Error: Se esperaba 1 fila afectada")
	}
*/
	return nil
}
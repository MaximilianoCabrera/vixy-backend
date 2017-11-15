package mysql

import (
	"log"
	"errors"
	"../../models"
	"../../dao/factory"
	"../../utilities"
)

type UsuarioImplMysql struct{}

func (dao UsuarioImplMysql) Create(u *models.Usuario) (string, error){
	query := "INSERT INTO usuario (nombre, apellido, nick, email, pass, idTipoUsuario, idImagen) VALUES (?, ?, ?, ?, ?, ?, ?)"
	db := get()
	defer db.Close()
	var msjError = ""

	//creo una sentencia=statement
	stmt, err := db.Prepare(query)
	if err != nil {
		msjError = "No se pudo crear la sentencia - stmt"
		return msjError, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.Nombre, u.Apellido, u.Nick, u.Email, u.Pass, 2, u.Imagen)
	if err != nil {
		msjError = "No se pudo guardar el usuario"
		return msjError, err
	}

	//con esto nos devuelve el ID
	id, err := result.LastInsertId()
	if err != nil {
		msjError = "No se pudo conseguir el ultimo id"
		return msjError, err
	}

	config, err := utilities.GetConfiguration()
	if err != nil{
		log.Fatalln(err)
	}
	//Guardo la imagen en la base de datos - tabla imagen
	imagenDAO := factory.ImagenFactoryDAO(config.Engine)
	var img models.Imagen
	img.Imagen = u.Imagen

	err = imagenDAO.Create(img)
	if err != nil{
		log.Fatalln(err)
	}

	u.ID = int(id)
	var msj = "Se pudo cargar el usuario y la imagen correctamente."
	return msj, nil
}
func (dao UsuarioImplMysql) GetAll() ([]models.Usuario, error) {
	query := "SELECT id, nombre, apellido, nick, email, pass, idTipoUsuario, idImagen FROM usuario"
	usuarios := make([]models.Usuario, 0)
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
		err := rows.Scan(&row.Nombre, &row.Apellido, &row.Nick, &row.Email, &row.IDTipoUsuario, &row.IDImagen)
		if err != nil {
			return usuarios, err
		}
		usuarios = append(usuarios, row)
	}
	return usuarios, nil
}
func (dao UsuarioImplMysql) GetByID(id int) (models.Usuario, error) {
	query := "SELECT id, nombre, apellido, nick, email, password, idTipoUsuario FROM usuario WHERE id = ?"

	db := get()
	defer db.Close()

	//creo una sentencia=statement
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Panic(err)
	}
	defer stmt.Close()
	row := db.QueryRow(query, id)

	var user = models.User{}

	err = row.Scan(&user.ID, &user.Nombre, &user.Apellido, &user.Nick, &user.Email, &user.Password, &user.IDTipoUsuario)
	if err != nil {
		return user, err
	}
	return user, nil
}
func (dao UsuarioImplMysql) GetOne(u models.Usuario) (models.Usuario, error){
}
func (dao UsuarioImplMysql) Update(u models.Usuario) error {
	query := "UPDATE usuario SET nombre = ?, apellido = ?, nick = ?, email = ?, password = ? WHERE id = ?"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row, err := stmt.Exec(u.Nombre, u.Apellido, u.Nick, u.Email, u.Password, u.ID)
	if err != nil {
		return err
	}

	i, _ := row.RowsAffected()
	if i != 1 {
		return errors.New("Error: Se esperaba 1 fila afectada")
	}
	return nil
}
func (dao UsuarioImplMysql) Delete(id int) error {
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
	return nil
}

package mysql

import (
	"log"
	"errors"
	"../../models"
	"../../dao/factory"
	"../../utilities"
)

type ImagenImplMysql struct{}

func (dao ImagenImplMysql) Create(i *models.Imagen) error {
	query := "INSERT INTO usuario (nombre, apellido, nick, email, pass, idTipoUsuario, idImagen) VALUES (?, ?, ?, ?, ?, ?, ?)"
	db := get()
	defer db.Close()

	//creo una sentencia=statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	config, err := utilities.GetConfiguration()
	if err != nil{
		log.Fatalln(err)
	}
	//TODO: Obtener el valor del id de la imagen - u.IDImagen
	tipoUsuarioDAO := factory.TipoUsuarioFactoryDAO(config.Engine)

	var tipoUsuario models.TipoUsuario
	tipoUsuario.Nombre = u.TipoUsuario

	tipoUsuarioDAO.GetOne(tipoUsuario)

	//imagenDAO := factory.ImagenFactoryDAO(config.Engine)


	result, err := stmt.Exec(u.Nombre, u.Apellido, u.Nick, u.Email, u.Pass, tipoUsuario, u.IDImagen)
	if err != nil {
		return err
	}

	//con esto nos devuelve el ID
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = int(id)
	return nil
}
func (dao ImagenImplMysql) GetAll() ([]models.Imagen, error) {
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
func (dao ImagenImplMysql) GetByID(id int) (models.Imagen, error) {
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
func (dao ImagenImplMysql) GetOne(i models.Imagen) (models.Imagen, error){
}
func (dao ImagenImplMysql) Update(i models.Imagen) error {
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
func (dao ImagenImplMysql) Delete(id int) error {
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

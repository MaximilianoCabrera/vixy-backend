package mysql

import (
	"../../models"
	"log"
	"errors"
)

type TipoUsuarioImplMysql struct{}

func (dao TipoUsuarioImplMysql) Create(tu models.TipoUsuario) error {
	query := "INSERT INTO tipoUsuario (id, nombre) VALUES (?, ?)"
	db := get()
	defer db.Close()

	//creo una sentencia=statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(tu.ID, tu.Nombre)
	if err != nil {
		return err
	}

	//con esto nos devuelve el ID
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	tu.ID = int(id)
	return nil
}
func (dao TipoUsuarioImplMysql) GetByID(id int) (models.TipoUsuario, error){
	query := "SELECT id, nombre FROM tipoUsuario WHERE id = ?"

	db := get()
	defer db.Close()

	//creo una sentencia=statement
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Panic(err)
	}
	defer stmt.Close()
	row := db.QueryRow(query, id)

	var tipoUsuario = models.TipoUsuario{}

	err = row.Scan(&tipoUsuario.ID, &tipoUsuario.Nombre)
	if err != nil {
		return tipoUsuario, err
	}

	return tipoUsuario, nil
}
func (dao TipoUsuarioImplMysql) GetAll() ([]models.TipoUsuario, error){
	query := "SELECT id, nombre FROM tipoUsuario"
	tipoUsuario := make([]models.TipoUsuario, 0)
	db := get()
	defer db.Close()

	//creo una sentencia=statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return tipoUsuario, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return tipoUsuario, err
	}

	for rows.Next() {
		var row models.TipoUsuario
		err := rows.Scan(&row.ID, &row.Nombre)
		if err != nil {
			return tipoUsuario, err
		}
		tipoUsuario = append(tipoUsuario, row)
	}
	return tipoUsuario, nil
}
func (dao TipoUsuarioImplMysql) Update(tu models.TipoUsuario) error{
	query := "UPDATE tipoUsuario SET nombre = ? WHERE id = ?"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row, err := stmt.Exec(tu.Nombre, tu.ID)
	if err != nil {
		return err
	}

	i, _ := row.RowsAffected()
	if i != 1 {
		return errors.New("Error: Se esperaba 1 fila afectada")
	}
	return nil
}
func (dao TipoUsuarioImplMysql) Delete(id int) error{
	query := "DELETE FROM users WHERE id = ?"
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
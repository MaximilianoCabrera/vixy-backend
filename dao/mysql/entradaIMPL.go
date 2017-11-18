package mysql

import (
	"../../models"
	//"log"
	//"errors"
)
type EntradaImplMysql struct{}

func (dao EntradaImplMysql ) Create (e *models.Entrada) error{
	/*
	query := "INSERT INTO entradas (titulo, imagen, url) VALUES (?, ?, ?)"
	db := DB()
	defer db.Close()

	//creo una sentencia=statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Titulo, e.Imagen, e.Url)
	if err != nil {
		return err
	}

	//con esto nos devuelve el ID
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	e.ID = int(id)
	*/
	return nil
}
func (dao EntradaImplMysql ) GetAll() ([]models.Entrada, error) {

//	query := "SELECT id, titulo, imagen, url FROM entradas"
	entradas := make([]models.Entrada, 0)
/*
	db := get()
	defer db.Close()

	//creo una sentencia=statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return entradas, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return entradas, err
	}

	for rows.Next() {
		var row models.Entrada
		err := rows.Scan(&row.ID, &row.Titulo, &row.Imagen, &row.Url)
		if err != nil {
			return entradas, err
		}
		entradas = append(entradas, row)
	}
	*/
	return entradas, nil
}
func (dao EntradaImplMysql ) GetByID(id int) (models.Entrada , error) {

	var entrada = models.Entrada{}
/*
	query := "SELECT id, titulo, imagen, url FROM entradas WHERE id = ?"

	db := get()
	defer db.Close()

	//creo una sentencia=statement
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	row := db.QueryRow(query, id)

	var entrada = models.Entrada{}

	err = row.Scan(&entrada.ID, &entrada.Titulo, &entrada.Imagen, &entrada.Url)

	if err != nil {
		return entrada , err
	}
*/
	return entrada , nil
}
func (dao EntradaImplMysql ) GetOne(e models.Entrada) (models.Entrada, error){
	//TODO: corregir para dejarlo en GETONE en vez de GETBYID
	var entrada = models.Entrada{}
/*
	query := "SELECT id, titulo, imagen, url FROM entradas WHERE id = ?"

	db := get()
	defer db.Close()

	//creo una sentencia=statement
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	row := db.QueryRow(query, e.ID)



	err = row.Scan(&entrada.ID, &entrada.Titulo, &entrada.Imagen, &entrada.Url)

	if err != nil {
		return entrada , err
	}
*/
	return entrada , nil
}
func (dao EntradaImplMysql ) Update(e models.Entrada) (models.Entrada, error) {
/*
	query := "UPDATE users SET titulo = ?, imagen = ?, url = ? WHERE id = ?"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return e,err
	}
	defer stmt.Close()

	row, err := stmt.Exec(e.Titulo, e.Imagen, e.Url, e.ID)
	if err != nil {
		return e, err
	}

	i, _ := row.RowsAffected()
	if i != 1 {
		return e, errors.New("Error: Se esperaba 1 fila afectada")
	}
*/
	return e, nil
}
func (dao EntradaImplMysql ) Delete(id int) error {
/*
	query := "DELETE FROM entradas WHERE id = ?"
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
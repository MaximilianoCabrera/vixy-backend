package mysql

import (

	"../../models"
	//"log"
	//"errors"
)

type ImagenImplMysql struct{}


func (dao ImagenImplMysql) Create(i *models.Imagen) (models.Imagen, error){
	var img models.Imagen
/*
	query := "INSERT INTO imagen (Imagen) VALUES (?)"
	db := get()
	defer db.Close()


	img.Imagen = i.Imagen

	//creo una sentencia=statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return img, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(i.Imagen)
	if err != nil {
		return img, err
	}

	//con esto nos devuelve el ID
	id, err := result.LastInsertId()
	if err != nil {
		return img, err
	}

	img.ID = int(id)
*/
	return img, nil
}
func (dao ImagenImplMysql) GetAll() ([]models.Imagen, error) {
	imgs := make([]models.Imagen, 0)
/*
	query := "SELECT id, imagen FROM imagen"

	db := get()
	defer db.Close()

	//creo una sentencia=statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return imgs, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return imgs, err
	}

	for rows.Next() {
		var row models.Imagen
		err := rows.Scan(&row.ID, &row.Imagen)
		if err != nil {
			return imgs, err
		}
		imgs = append(imgs, row)
	}
*/
	return imgs, nil
}
func (dao ImagenImplMysql) GetByID(id int) (models.Imagen, error) {
	var img = models.Imagen{}
/*
	query := "SELECT imagen FROM imagen WHERE id = ?"

	db := get()
	defer db.Close()

	//creo una sentencia=statement
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	row := db.QueryRow(query, id)

	err = row.Scan(&img.ID, &img.Imagen)
	if err != nil {
		return img, err
	}
*/
	return img, nil
}
func (dao ImagenImplMysql) GetOne(i models.Imagen) (models.Imagen, error){
	var imagen = models.Imagen{}
/*
	query := "SELECT id, imagen FROM imagen WHERE id = ?"

	db := get()
	defer db.Close()

	//creo una sentencia=statement
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	row := db.QueryRow(query, i.ID, i.Imagen)



	err = row.Scan(&imagen.ID, &imagen.Imagen)
	if err != nil {
		return imagen, err
	}

*/
	return imagen, nil
}
func (dao ImagenImplMysql) Update(i models.Imagen) (models.Imagen, error) {
	var img models.Imagen
	img.Imagen = i.Imagen

/*
	query := "UPDATE imagen SET nombre = ? WHERE id = ?"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return i, err
	}
	defer stmt.Close()


	//var imgVieja models.Imagen



	row, err := stmt.Exec(i.Imagen)
	if err != nil {
		return i, err
	}

	e, _ := row.RowsAffected()
	if e != 1 {
		return i, errors.New("Error: Se esperaba 1 fila afectada")
	}
*/
	return i, nil
}
func (dao ImagenImplMysql) Delete(id int) error {
	/*
	query := "DELETE FROM imagen WHERE id = ?"
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

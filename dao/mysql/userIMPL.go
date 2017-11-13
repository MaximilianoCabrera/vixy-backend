package mysql

import (
	"errors"
	"../../models"
	"log"
)

type UserImplMysql struct{}

func (dao UserImplMysql) Create(u models.User) error {
	query := "INSERT INTO users (nombre, apellido, nick, email, password) VALUES (?, ?, ?, ?, ?)"
	db := get()
	defer db.Close()

	//creo una sentencia=statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.Nombre, u.Apellido, u.Nick, u.Email, u.Password)
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
func (dao UserImplMysql) GetByID(id int) (models.User, error) {
	query := "SELECT id, nombre, apellido, nick, email, apellido FROM users WHERE id = ?"

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

	err = row.Scan(&user.ID, &user.Nombre, &user.Nick, &user.Apellido, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}
func (dao UserImplMysql) GetAll() ([]models.User, error) {
	query := "SELECT id, nombre, apellido, nick, email, password FROM users"
	users := make([]models.User, 0)
	db := get()
	defer db.Close()

	//creo una sentencia=statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return users, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var row models.User
		err := rows.Scan(&row.ID, &row.Nombre, &row.Apellido, &row.Nick, &row.Email, &row.Password)
		if err != nil {
			return users, err
		}
		users = append(users, row)
	}
	return users, nil
}
func (dao UserImplMysql) Update(u models.User) error {
	query := "UPDATE users SET nombre = ?, apellido = ?, nick = ?, email = ?, password = ? WHERE id = ?"
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
func (dao UserImplMysql) Delete(id int) error {
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

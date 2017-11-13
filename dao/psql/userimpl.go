package psql

import (
	"errors"
	"../../models"
)

type UserImplPsql struct{}

func (dao UserImplPsql) Create(u *models.User) error {
	query := "INSERT INTO users (nombre, apellido, email) VALUES ($1, $2, $3) RETURNING id"
	db := get()
	defer db.Close()

	//creo una sentencia=statement
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row := stmt.QueryRow(u.Nombre, u.Apellido, u.Email)
	row.Scan(&u.ID)
	return nil
}

func (dato UserImplPsql) GetAll() ([]models.User, error) {
	query := "SELECT id, nombre, apellido, email FROM users"
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
		err := rows.Scan(&row.ID, &row.Nombre, &row.Apellido, &row.Email)
		if err != nil {
			return users, err
		}
		users = append(users, row)
	}
	return users, nil
}

func (dao UserImplPsql) Update(u *models.User) error {
	query := "UPDATE users SET nombre = $1, apellido = $2, email = $3 WHERE id = $4"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row, err := stmt.Exec(u.Nombre, u.Apellido, u.Email, u.ID)
	if err != nil {
		return err
	}

	i, _ := row.RowsAffected()
	if i != 1 {
		return errors.New("Error: Se esperaba 1 fila afectada")
	}
	return nil
}

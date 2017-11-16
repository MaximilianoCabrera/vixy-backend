package mysql

import (
	"log"
	"errors"
	"../../models"
)

type UsuarioImplMysql struct{}
//TODO: Hacer que le create reciba los datos de la imagen a cargar, para así poder hacer
//TODO: Una verdadera transacción, controlando que se cargue todo, o hacer Rollback.
//OK//
func (dao UsuarioImplMysql) Create(u *models.Usuario) (string, error){
	db := get()
	defer db.Close()

	db.Begin()

	query := "INSERT INTO usuario (nombre, apellido, nick, email, password, idTipoUsuario, idImagen) VALUES (?, ?, ?, ?, ?, ?, ?)"

	var msjError = ""

	//creo una sentencia=statement
	stmt, err := db.Prepare(query)
	if err != nil {
		msjError = "No se pudo crear la sentencia - stmt"
		return msjError, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.Nombre, u.Apellido, u.Nick, u.Email, u.Pass, 2, u.IDImagen)

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

	var img models.Imagen
	img.ID = u.IDImagen

	u.ID = int(id)
	var msj = "Se pudo cargar el usuario %s, y la imagen %i correctamente."
	return msj, nil
}
func (dao UsuarioImplMysql) GetAll() ([]models.Usuario, error) {
	query := "SELECT * FROM usuario"
	var usuarios []models.Usuario
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
	return usuarios, nil
}

//Probar//
func (dao UsuarioImplMysql) GetByID(id int) (models.Usuario, error) {
	query := "SELECT * FROM usuario WHERE id = ?"

	db := get()
	defer db.Close()

	//creo una sentencia=statement
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Panic(err)
	}
	defer stmt.Close()
	row := db.QueryRow(query, id)

	var user = models.Usuario{}

	//TODO: obtener y transformar el tipoUsuario
	err = row.Scan(&user.ID, &user.Nombre, &user.Apellido, &user.Nick, &user.Email, &user.Pass, &user.TipoUsuario, &user.IDImagen)
	if err != nil {
		log.Println("Error: ", err)
		return user, err
	}
	log.Println("Sale bien")
	return user, nil
}
func (dao UsuarioImplMysql) GetOne(u models.Usuario) (models.Usuario, error){
	query := "SELECT id, nombre, apellido, nick, email, password, idTipoUsuario FROM usuario " +
		"WHERE id = ? " +
			"OR nombre = ? " +
			"OR apellido = ? " +
			"OR nick = ? " +
			"OR email = ? " +
			"OR password = ? " +
			"OR idTipoUsuario = ? "

	db := get()
	defer db.Close()

	//creo una sentencia=statement
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Panic(err)
	}
	defer stmt.Close()

/*
	//Obtengo el id del TipoUsuario
	config, err := utilities.GetConfiguration()
	if err != nil{
		log.Fatalln(err)
	}
	var tipoUsuario models.TipoUsuario
	tipoUsuario.Nombre = u.TipoUsuario

	if tipoUsuario.Nombre != ""{
		tipoUsuarioDAO := factory.TipoUsuarioFactoryDAO(config.Engine)
		tipoUsuario, err = tipoUsuarioDAO.GetOne(tipoUsuario)
	}
*/
	//row := db.QueryRow(query, u.ID, u.Nombre, u.Apellido, u.Nick, u.Email, u.Pass, tipoUsuario, )

	var usuario = models.Usuario{}

	//err = row.Scan(&usuario.ID, &usuario.Nombre, &usuario.Apellido, &usuario.Nick, &usuario.Email, &usuario.Pass, &usuario.TipoUsuario)
	if err != nil {
		return usuario, err
	}
	return usuario, nil
}
func (dao UsuarioImplMysql) Update(u models.Usuario) (models.Usuario, error) {
	query := "UPDATE usuario SET nombre = ?, apellido = ?, nick = ?, email = ?, password = ?, idImagen = ? WHERE id = ?"
	db := get()
	defer db.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return u ,err
	}
	defer stmt.Close()

	//TODO: Corregir aca.
	var img models.Imagen
	img.ID = u.IDImagen
/*
	if u.Imagen != ""{
		config, err := utilities.GetConfiguration()
		if err != nil{
			log.Fatalln(err)
		}
		//Guardo la imagen en la base de datos - tabla imagen
		imagenDAO := factory.ImagenFactoryDAO(config.Engine)

		img, err = imagenDAO.Create(img)
		if err != nil{
			log.Fatalln(err)
		}
	}
*/
	row, err := stmt.Exec(u.Nombre, u.Apellido, u.Nick, u.Email, u.Pass, img.ID, u.ID)
	if err != nil {
		return u, err
	}

	i, _ := row.RowsAffected()
	if i != 1 {
		return u, errors.New("Error: Se esperaba 1 fila afectada con los siguientes datos")
	}
	return u, nil
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
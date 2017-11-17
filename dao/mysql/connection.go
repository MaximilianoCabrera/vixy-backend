package mysql

import (
	"database/sql"
	"fmt"
	"../../utilities"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"github.com/azer/crud"
)

func DB() *crud.DB{
	var DB *crud.DB
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true", config.User, config.Password, config.Server, config.Port, config.Database)
	DB, err = crud.Connect("mysql", dataSourceName)
	if err != nil{
		fmt.Println("Error en DB: ", err)
	}

	err = DB.Ping()
	fmt.Println(DB.Ping())

	return DB
}

func get() *sql.DB {
		config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true", config.User, config.Password, config.Server, config.Port, config.Database)

	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

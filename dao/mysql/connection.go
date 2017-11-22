package mysql

import (
	"log"
	"fmt"

	"../../utilities"
	"github.com/azer/crud"
	_ "github.com/go-sql-driver/mysql"
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
		log.Fatalln("Error en DB: ", err)
		//fmt.Println("Error en DB: ", err)
	}

	err = DB.Ping()

	return DB
}
package psql

import (
	"database/sql"
	"../../utilities"
	"log"

	"fmt"
	_ "github.com/lib/pq"
)

func get() *sql.DB {
	config, err := utilities.GetConfiguration()
	if err != nil {
		log.Fatalln(err)
	}

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.User, config.Password, config.Port, config.Database)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

package actions

import (
	"net/http"
	"log"
)

func Close(w http.ResponseWriter, r *http.Request){
	log.Fatalln("Cerrando APP")
}
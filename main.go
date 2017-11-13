package main

import (
	"net/http"
	"./routes"
	"github.com/rs/cors"
)


func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
		AllowCredentials: true,
	})

	router := routes.NewMainRouter()

	http.ListenAndServe(":8080", c.Handler(router))
}

package main

import (
	"log"
	"net/http"

	"github.com/s0ld13r/golang-microservices/introduction/mvc/app"
)

func main() {
	app.StartApp()
	log.Println("Starting server on port", 8080)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

package app

import (
	"net/http"

	"github.com/s0ld13r/golang-microservices/introduction/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)
}

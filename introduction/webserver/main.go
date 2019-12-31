package main

import (
	"log"
	"net/http"
)

const port = 8081

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		if _, err := writer.Write([]byte("Hello world")); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}
	})
	log.Println("Listening on port", port)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}

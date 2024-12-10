package main

import (
	"bookapi/handlers"
	"log"
	"net/http"
)

func main() {

	mux := handlers.NewHandler()
	log.Fatal(http.ListenAndServe("localhost:5000", mux))

}

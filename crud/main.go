package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	fmt.Println("Servidor up em localhost:5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}

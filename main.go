package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/goriila/mux"
)

func testRoute(w http.ResponseWriter, r *http.Request) {

	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}

func main() {

	fmt.Println("This is a test")

}

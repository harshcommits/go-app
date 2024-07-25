package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func testRoute(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Testing from backend")

}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", testRoute)
	fmt.Println("Server started and listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

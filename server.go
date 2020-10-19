package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/bstick12/goflake"
	"github.com/gorilla/mux"
)

var generator = goflake.GoFlakeInstanceUsingUnique("D01Z01")

func main() {
	startServer()
}

func startServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HelloWorld)
	router.Queries("count", "{count:[0-9]+}")
	log.Println("Hello world!")
	log.Println("Starting server...")
	log.Println("Navigate to localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// HelloWorld prints a Hello World statement
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello, KubeCon!")

}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func get_health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Println("health end point was hit")
}

func handle_requests() {
	r := mux.NewRouter()
	r.HandleFunc("/get_health", get_health).Methods(http.MethodGet)
	log.Fatalln(http.ListenAndServe(":8080", r))
}

func main() {
	fmt.Println("running")
	handle_requests()
	fmt.Println("maybe not running")
}

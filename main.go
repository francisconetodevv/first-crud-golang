package main

import (
	"CRUD/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Vai conter toda as rotas para interagir com o banco
	router := mux.NewRouter()

	// Routers with the CRUD actions
	router.HandleFunc("/user", server.CreateUser).Methods(http.MethodPost)        // Create
	router.HandleFunc("/users", server.SearchUsers).Methods(http.MethodGet)       // Read - Search
	router.HandleFunc("/user/{id}", server.SearchUser).Methods(http.MethodGet)    // Read - Search
	router.HandleFunc("/user/{id}", server.UpdateUser).Methods(http.MethodPut)    // Update
	router.HandleFunc("/user/{id}", server.DeleteUser).Methods(http.MethodDelete) // Delete

	fmt.Println("Escutando na rota 5000")
	log.Fatal(http.ListenAndServe(":5000", router))

}

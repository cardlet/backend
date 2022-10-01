package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func createRouter() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeLink)

	// Authenticated
	router.HandleFunc("/user/login", loginUser).Methods("POST")
	router.HandleFunc("/user/register", registerUser).Methods("POST")
	router.HandleFunc("/user/update", updateUser).Methods("PATCH")
	router.HandleFunc("/user/delete", deleteUser).Methods("DELETE")
	router.HandleFunc("/user/desk/create", createDesk).Methods("POST")

	// Public
	router.HandleFunc("/users", getAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getOneUser).Methods("GET")

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8000", handler))
}
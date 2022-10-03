package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func createRouter() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeLink)

	// User routes
	router.HandleFunc("/user/login", loginUser).Methods("POST")
	router.HandleFunc("/user/register", registerUser).Methods("POST")
	router.HandleFunc("/user/update", updateUser).Methods("PATCH")
	router.HandleFunc("/user/delete", deleteUser).Methods("DELETE")

	// Desk routes
	router.HandleFunc("/desk/create", createDesk).Methods("POST")
	router.HandleFunc("/desks", getAllDesks).Methods("GET")
	router.HandleFunc("/desks/{id}", getDesksByUser).Methods("GET")
	router.HandleFunc("/desk/insert", insertCard).Methods("POST")

	// Public routes
	router.HandleFunc("/users", getAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getOneUser).Methods("GET")

	router.Use(contentTypeApplicationJsonMiddleware)

	handler := cors.Default().Handler(router)

	port := config["SERVER_PORT"].String()

	fmt.Println("Running at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":" + port, handler))
}
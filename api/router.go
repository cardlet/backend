package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	createMessageResponse(w, "API is online")
}

func CreateRouter() {
	router := chi.NewRouter()

	router.Use(contentTypeApplicationJsonMiddleware)

	router.HandleFunc("/", homeLink)

	// User routes
	router.Post("/user/login", loginUser)
	router.Post("/user/register", registerUser)
	router.Patch("/user/update", updateUser)
	router.Delete("/user/delete", deleteUser)

	// Desk routes
	router.Post("/desk/create", createDesk)
	router.Get("/desks", getAllDesks)
	router.Get("/desks/{id}", getDesksByUser)
	router.Post("/desk/insert", insertCard)

	// Public routes
	router.Get("/users", getAllUsers)
	router.Get("/users/{id}", getOneUser)

	handler := cors.Default().Handler(router)

	port := config["SERVER_PORT"].String()

	fmt.Println("Running at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":" + port, handler))
}
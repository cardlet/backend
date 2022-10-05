package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	createMessageResponse(w, "API is online")
}

func CreateRouter() {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	  }))

	router.Use(contentTypeApplicationJsonMiddleware)

	router.HandleFunc("/", homeLink)

	// User routes
	router.Post("/user/login", loginUser)
	router.Post("/user/register", registerUser)
	router.Patch("/user/update", updateUser)
	router.Delete("/user/delete", deleteUser)

	// Public routes
	router.Get("/users", getAllUsers)
	router.Get("/users/{id}", getOneUser)

	// Desk routes
	router.Post("/desk/create", createDesk)
	router.Get("/desks", getAllDesks)
	router.Get("/desks/{id}", getDesksByUser)

	// Card routes
	router.Post("/card/create", createCard)
	router.Get("/cards", getAllCards)
	router.Get("/cards/{id}", getCardsByUser)
	router.Get("/user/cards/{id}", getCardsByDeskId)

	port := config["SERVER_PORT"].String()

	fmt.Println("Running at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":" + port, router))
}
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
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Use(contentTypeApplicationJsonMiddleware)

	router.HandleFunc("/", homeLink)

	// User routes
	router.Route("/user", UserRouter)

	// Public routes
	router.Route("/users", PublicUserRouter)

	// Desk routes
	router.Route("/desks", DesksRouter)

	// Card routes
	router.Route("/cards", CardsRouter)

	port := config["SERVER_PORT"].String()

	fmt.Println("Running at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func UserRouter(r chi.Router) {
	r.Post("/login", loginUser)
	r.Post("/register", registerUser)
	r.Patch("/update", updateUser)
	r.Delete("/delete", deleteUser)
}

func PublicUserRouter(r chi.Router) {
	r.Get("/", getAllUsers)
	r.Get("/{id}", getOneUser)
}

func DesksRouter(r chi.Router) {
	r.Get("/", getAllDesks)
	r.Post("/create", createDesk)
	r.Get("/{userId}", getDesksByUser)
	r.Delete("/delete/{id}", deleteDesk)
}

func CardsRouter(r chi.Router) {
	r.Get("/", getAllCards)
	r.Post("/create", createCard)
	r.Get("/user", getCardsByUser)
	r.Get("/{id}", getCardsByDeskId)
	r.Delete("/delete/{id}", deleteCard)
}

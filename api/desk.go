package api

import (
	"net/http"
	"strconv"

	"github.com/cardlet/obj"
	"github.com/go-chi/chi/v5"
)

func createDesk(w http.ResponseWriter, r *http.Request) {
	user, ok := validateUser(r)
	if !ok {
		return
	}

	var desk obj.Desk

	if !UnmarshalJsonBody(w, r, desk) {
		return
	}

	desk.UserID = user.ID
	db.Create(&desk)

	w.WriteHeader(http.StatusCreated)
}

func getAllDesks(w http.ResponseWriter, r *http.Request) {
	var desks []obj.Desk
	db.Find(&desks)
	createJsonResponse(w, desks)
}

func getDesksByUser(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)

	var desks []obj.Desk
	sampleDesk := obj.Desk{
		UserID: uint(userId),
	}
	db.Find(&desks, &sampleDesk)

	createJsonResponse(w, desks)
}


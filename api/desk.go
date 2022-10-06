package api

import (
	"net/http"

	"github.com/cardlet/obj"
)

func createDesk(w http.ResponseWriter, r *http.Request) {
	user, ok := validateUser(w, r)
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
	userId := getUintParam(r, "userId")

	var desks []obj.Desk
	sampleDesk := obj.Desk{
		UserID: uint(userId),
	}
	db.Find(&desks, &sampleDesk)

	createJsonResponse(w, desks)
}


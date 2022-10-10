package api

import (
	"net/http"
	"fmt"
	"github.com/cardlet/obj"
)

func createDesk(w http.ResponseWriter, r *http.Request) {
	user, ok := validateUser(w, r)
	if !ok {
		return
	}

	var desk obj.Desk

	if !UnmarshalJsonBody(w, r, &desk) {
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

func deleteDesk(w http.ResponseWriter, r *http.Request) {
	user, ok := validateUser(w, r)
	if !ok {
		return
	}

	var desk obj.Desk
	deskId := getUintParam(r, "id")
	db.Find(&desk, "ID = ?", deskId)

	if (desk.UserID != user.ID) {
		createErrorResponse(w, "No permissions to delete the desk!")
		return
	}

	db.Delete(&desk)
	createMessageResponse("Successfully deleted the desk!")
}
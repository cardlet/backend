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

	if !UnmarshalJsonBody(w, r, &desk) {
		return
	}

	desk.UserID = user.ID
	db.Create(&desk)

	w.WriteHeader(http.StatusCreated)

	createJsonResponse(w, desk)
}

func getAllDesks(w http.ResponseWriter, r *http.Request) {
	var desks []obj.Desk
	db.Find(&desks)

	desks = DesksPlusCardCount(desks)

	createJsonResponse(w, desks)
}

func getDesksByUser(w http.ResponseWriter, r *http.Request) {
	user, ok := validateUser(w, r)

	if !ok {
		return
	}

	var desks []obj.Desk
	sampleDesk := obj.Desk{
		UserID: uint(user.ID),
	}
	db.Find(&desks, &sampleDesk)

	desks = DesksPlusCardCount(desks)

	createJsonResponse(w, desks)
}

func deleteDesk(w http.ResponseWriter, r *http.Request) {
	user, ok := validateUser(w, r)
	if !ok {
		return
	}

	var desk obj.Desk
	sampleDesk := obj.Desk{}
	sampleDesk.ID = uint(getUintParam(r, "id"))

	db.Find(&desk, sampleDesk)

	if desk.UserID != user.ID {
		createErrorResponse(w, "No permissions to delete the desk!")
		return
	}

	db.Delete(&desk)
	createMessageResponse(w, "Successfully deleted the desk!")
}

func DesksPlusCardCount(selectedDesks []obj.Desk) []obj.Desk {
	desks := selectedDesks
	for i := range desks {
		var cards []obj.Card
		sampleCard := obj.Card{DeskID: desks[i].ID}
		db.Find(&cards, sampleCard)
		desks[i].CardCount = uint(len(cards))
	}
	return desks
}

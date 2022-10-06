package api

import (
	"net/http"

	"github.com/cardlet/obj"
)

func createCard(w http.ResponseWriter, r *http.Request) {
	user, ok := validateUser(w, r)
	if !ok {
		return
	}

	var card obj.Card
	if !UnmarshalJsonBody(w, r, &card) {
		return
	}

	card.UserID = user.ID

	db.Create(&card)
}

func getAllCards(w http.ResponseWriter, r *http.Request) {
	var cards []obj.Card
	db.Find(&cards)
	createJsonResponse(w, cards)
}

func getCardsByUser(w http.ResponseWriter, r *http.Request) {
	userId := getUintParam(r, "userId")

	var cards []obj.Card
	sampleDesk := obj.Card{
		UserID: uint(userId),
	}
	db.Find(&cards, &sampleDesk)

	createJsonResponse(w, cards)
}

func getCardsByDeskId(w http.ResponseWriter, r *http.Request) {
	deskId := getUintParam(r, "deskId")
	
	var cards []obj.Card
	sampleDesk := obj.Card{
		DeskID: uint(deskId),
	}
	db.Find(&cards, &sampleDesk)

	createJsonResponse(w, cards)
}
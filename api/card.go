package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/cardlet/obj"
	"github.com/go-chi/chi/v5"
)

func createCard(w http.ResponseWriter, r *http.Request) {
	user, ok := validateUser(r)
	if !ok {
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the user title and description only in order to update")
	}
	var card obj.Card
	json.Unmarshal(reqBody, &card)
	card.UserID = user.ID

	db.Create(&card)
}

func getAllCards(w http.ResponseWriter, r *http.Request) {
	var cards []obj.Card
	db.Find(&cards)
	createJsonResponse(w, cards)
}

func getCardsByUser(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.ParseUint(chi.URLParam(r, "id"), 10, 32)

	var cards []obj.Card
	sampleDesk := obj.Card{
		UserID: uint(userId),
	}
	db.Find(&cards, &sampleDesk)

	createJsonResponse(w, cards)
}

func getCardsByDeskId(w http.ResponseWriter, r *http.Request) {
	deskId, _ := strconv.ParseUint(chi.URLParam(r, "deskId"), 10, 32)
	
	var cards []obj.Card
	sampleDesk := obj.Card{
		DeskID: uint(deskId),
	}
	db.Find(&cards, &sampleDesk)

	createJsonResponse(w, cards)
}
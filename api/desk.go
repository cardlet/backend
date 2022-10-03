package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/cardlet/obj"
	"github.com/gorilla/mux"
)

func createDesk(w http.ResponseWriter, r *http.Request) {
	user, ok := validateUser(r)
	if !ok {
		return
	}

	var desk obj.Desk

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		createErrorResponse(w, "Kindly enter data with the user title and description only in order to update")
	}

	json.Unmarshal(reqBody, &desk)
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
	userId, _ := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)

	var desks []obj.Desk
	sampleDesk := obj.Desk{
		UserID: uint(userId),
	}
	db.Find(&desks, &sampleDesk)

	createJsonResponse(w, desks)
}

func insertCard(w http.ResponseWriter, r *http.Request) {
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
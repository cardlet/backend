package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func createDesk(w http.ResponseWriter, r *http.Request) {
	user, ok := validateUser(r)
	if !ok {
		return
	}

	var desk Desk

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the user title and description only in order to update")
	}

	json.Unmarshal(reqBody, &desk)
	desk.UserID = user.ID
	db.Create(&desk)

	w.WriteHeader(http.StatusCreated)
}



func getAllDesks(w http.ResponseWriter, r *http.Request) {
	var desks []Desk
	db.Find(&desks)
	json.NewEncoder(w).Encode(desks)
}

func getDesksByUser(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)

	var desks []Desk
	sampleDesk := Desk{
		UserID: uint(userId),
	}
	db.Find(&desks, &sampleDesk)

	json.NewEncoder(w).Encode(desks)
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
	var card Card
	json.Unmarshal(reqBody, &card)
	card.UserID = user.ID

	db.Create(&card)
}
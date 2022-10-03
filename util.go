package main

import (
	"encoding/json"
	"net/http"
)

func createJsonResponse(w http.ResponseWriter, resp any) {
	json.NewEncoder(w).Encode(resp)
}

func createMessageResponse(w http.ResponseWriter, msg string) {
	json.NewEncoder(w).Encode(
		Message {
			Message: msg,
		},
	)
}
package main

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
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

func createErrorResponse(w http.ResponseWriter, msg string) {
	json.NewEncoder(w).Encode(
		ErrorMessage {
			Error: msg,
		},
	)
}

func createTokenResponse(w http.ResponseWriter, token string) {
	json.NewEncoder(w).Encode(
		TokenResponse {
			Token: token,
		},
	)
	w.WriteHeader(http.StatusOK)
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
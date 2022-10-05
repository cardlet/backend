package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/cardlet/obj"
	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

func createJsonResponse(w http.ResponseWriter, resp any) {
	json.NewEncoder(w).Encode(resp)
}

func createMessageResponse(w http.ResponseWriter, msg string) {
	json.NewEncoder(w).Encode(
		obj.Message {
			Message: msg,
		},
	)
}

func createErrorResponse(w http.ResponseWriter, msg string) {
	json.NewEncoder(w).Encode(
		obj.ErrorMessage {
			Error: msg,
		},
	)
}

func createTokenResponse(w http.ResponseWriter, token string) {
	json.NewEncoder(w).Encode(
		obj.TokenResponse {
			Token: token,
		},
	)
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func UnmarshalJsonBody(w http.ResponseWriter, r *http.Request, obj interface{}) bool {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		createErrorResponse(w, "Invalid response body")
		return false
	}
	err = json.Unmarshal(reqBody, obj)
	if err != nil {
		createErrorResponse(w, "Invalid request")
		return false
	}
	return true
}

func getUintParam(r *http.Request, key string) uint64 {
	num, err := strconv.ParseUint(chi.URLParam(r, "deskId"), 10, 32)
	if err != nil {
		return 0
	}
	return num
}
package main

import (
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

func getOneUser(w http.ResponseWriter, r *http.Request) {
	var user User
	db.Find(&user, "ID = ?", mux.Vars(r)["id"])
	json.NewEncoder(w).Encode(user)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []User
	db.Find(&users)
	createJsonResponse(w, users)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	user, ok := validateUser(r)
	if !ok {
		return
	}

	var updatedUser User

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		createErrorResponse(w, "Kindly enter data with the user title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedUser)

	user.Bio = updatedUser.Bio
	user.Name = updatedUser.Name
	user.Pass = updatedUser.Pass
	
	db.Model(&user).Updates(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	user, ok := validateUser(r)
	if !ok {
		return
	}
	db.Delete(&user, "Token = ?", user.Token)
}

func GenerateSecureToken(length int) string {
    b := make([]byte, length)
    if _, err := rand.Read(b); err != nil {
        return ""
    }
    return hex.EncodeToString(b)
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		createErrorResponse(w, "Kindly enter data with the user title and description only in order to update")
	}
	
	json.Unmarshal(reqBody, &newUser)

	var nameUser User
	result := db.First(&nameUser, "Name = ?", newUser.Name)

	// Name already exists
	if result.Error == nil {
		createErrorResponse(w, "Username already taken!")
	}

	if newUser.Pass == "" {
		createErrorResponse(w, "Invalid password!")
		return
	}

	newUser.Pass, _ = HashPassword(newUser.Pass)

	newUser.Token = GenerateSecureToken(20)
	db.Create(&newUser)

	w.WriteHeader(http.StatusCreated)

	createTokenResponse(w, newUser.Token)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	var user User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		createErrorResponse(w, "Kindly enter data with the user title and description only in order to update")
	}
	
	json.Unmarshal(reqBody, &user)

	var dbUser User
	db.First(&dbUser, "Name = ?", user.Name)

	if !CheckPasswordHash(user.Pass, dbUser.Pass) {
		createErrorResponse(w, "Access denied")
		return
	}
	createTokenResponse(w, dbUser.Token)
}

func validateUser(r *http.Request) (*User, bool) {
	var token = r.Header.Get("x-access-token")
	var user User
	db.First(&user, "Token = ?", token)
	if &user != nil {
		return &user, true
	}
	return nil, false
}
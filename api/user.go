package api

import (
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/cardlet/obj"
	"github.com/gorilla/mux"
)

func getOneUser(w http.ResponseWriter, r *http.Request) {
	var user obj.User
	db.Find(&user, "ID = ?", mux.Vars(r)["id"])
	json.NewEncoder(w).Encode(user)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []obj.User
	db.Find(&users)
	createJsonResponse(w, users)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	user, ok := validateUser(r)
	if !ok {
		return
	}

	var updatedUser obj.User

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
	var newUser obj.User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		createErrorResponse(w, "Kindly enter data with the user title and description only in order to update")
	}
	
	json.Unmarshal(reqBody, &newUser)

	var nameTaken bool
	db.Model(&obj.User{}).
	Select("count(*) > 0").
	Select("Name = ?", newUser.Name).
	Find(&nameTaken)

	// Name already exists
	if nameTaken {
		createErrorResponse(w, "Username already taken!")
		return
	}

	if newUser.Pass == "" {
		createErrorResponse(w, "Invalid password!")
		return
	}

	newUser.Pass, _ = HashPassword(newUser.Pass)

	newUser.Token = GenerateSecureToken(20)
	db.Create(&newUser)

	createTokenResponse(w, newUser.Token)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	var user obj.User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		createErrorResponse(w, "Kindly enter data with the user title and description only in order to update")
	}
	
	json.Unmarshal(reqBody, &user)

	var nameTaken bool
	db.Model(&obj.User{}).
	Select("count(*) > 0").
	Select("Name = ?", user.Name).
	Find(&nameTaken)

	if !nameTaken {
		createErrorResponse(w, "User doesn't exist!")
		return
	}

	var dbUser obj.User
	db.First(&dbUser, "Name = ?", user.Name)

	if !CheckPasswordHash(user.Pass, dbUser.Pass) {
		createErrorResponse(w, "Access denied")
		return
	}
	createTokenResponse(w, dbUser.Token)
}

func validateUser(r *http.Request) (*obj.User, bool) {
	var token = r.Header.Get("x-access-token")
	var user obj.User
	db.First(&user, "Token = ?", token)
	if &user != nil {
		return &user, true
	}
	return nil, false
}
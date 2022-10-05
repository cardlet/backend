package api

import (
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/cardlet/obj"
	"github.com/go-chi/chi/v5"
)

func getOneUser(w http.ResponseWriter, r *http.Request) {
	var user obj.User
	db.Find(&user, "ID = ?", chi.URLParam(r, "id"))
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
	if !UnmarshalJsonBody(w, r, &updatedUser) {
		return
	}

	user.Bio = updatedUser.Bio
	user.Name = updatedUser.Name
	user.Pass = updatedUser.Pass
	
	db.Model(&user).Updates(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	user, ok := validateUser(r)
	if !ok {
		createErrorResponse(w, "Authorization failed!")
		return
	}

	createMessageResponse(w, "Success!")
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
	if !UnmarshalJsonBody(w, r, &newUser) {
		return
	}

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
	if !UnmarshalJsonBody(w, r, &user) {
		return
	}

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
	var token = r.Header.Get("Authorization")

	var user obj.User
	err := db.First(&user, "Token = ?", token).Error

	if err == nil {
		return &user, true
	}
	return nil, false
}
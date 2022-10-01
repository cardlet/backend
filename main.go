package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var users = []User{
	{
		Name: "First User",
		Bio: "Welcome to my profile",
		Token: "1",
	},
}

var db *gorm.DB

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API is online")
}

func getOneUser(w http.ResponseWriter, r *http.Request) {
	for _, singleUser := range users {
		if singleUser.Token == mux.Vars(r)["token"] {
			json.NewEncoder(w).Encode(singleUser)
		}
	}
}

func getAllFromDb() []User {
	var users []User
	db.Find(&users)
	return users
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(getAllFromDb())
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	i, singleUser, ok := validateUser(r)
	if !ok {
		return
	}

	var updatedUser User

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the user title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedUser)
	
	singleUser.Name = updatedUser.Name
	singleUser.Bio = updatedUser.Bio
	users = append(users[:i], *singleUser)
	json.NewEncoder(w).Encode(singleUser)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	index, _, ok := validateUser(r)
	if !ok {
		return
	}
	users = append(users[:index], users[index+1:]...)
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
		fmt.Fprintf(w, "Kindly enter data with the user title and description only in order to update")
	}
	
	json.Unmarshal(reqBody, &newUser)
	newUser.Token = GenerateSecureToken(20)
	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser.Token)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	_, user, ok := validateUser(r)
	if !ok {
		fmt.Fprintf(w, "Access denied")
		return
	}
	json.NewEncoder(w).Encode(user)
}

func validateUser(r *http.Request) (int, *User, bool) {
	var token = r.Header.Get("x-access-token")
	for i, singleUser := range users {
		if singleUser.Token == token {
			return i, &singleUser, true
		}
	}
	return 0, nil, false
}

func createDesk(w http.ResponseWriter, r *http.Request) {
	index, user, ok := validateUser(r)
	if !ok {
		return
	}

	var desk Desk

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the user title and description only in order to update")
	}

	json.Unmarshal(reqBody, &desk)
	user.Desks = append(user.Desks, desk)
	users = append(users[:index], *user)
	w.WriteHeader(http.StatusCreated)
}

func main() {
	var err error
	db, err = gorm.Open(postgres.Open("postgres://postgres@localhost/test"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	  }
	
	db.AutoMigrate(&User{})	
	db.Create(&User{
		Name: "Zweiter",
		Bio:  "Ne",
		Pass: "",
		Token: GenerateSecureToken(20),
	})

	createRouter()
}
package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name 	string `json:"name"`
	Bio 	string `json:"bio"`
	Pass	string `json:"password"`
	Token	string `json:"token"`
	Desks	[]Desk `json:"desks"`
}

type Desk struct {
	Name	string `json:"name"`
	Cards	[]Card `json:"cards"`
}

type Card struct {
	Quest	string	`json:"quest"`
	Answer	string	`json:"answer"`
}

var users = []User{
	{
		Name: "First User",
		Bio: "Welcome to my profile",
		Token: "1",
	},
}

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

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
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

	json.NewEncoder(w).Encode(newUser)
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
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homeLink)

	// Authenticated
	router.HandleFunc("/user/login", loginUser).Methods("POST")
	router.HandleFunc("/user/register", registerUser).Methods("POST")
	router.HandleFunc("/user/update", updateUser).Methods("PATCH")
	router.HandleFunc("/user/delete", deleteUser).Methods("DELETE")
	router.HandleFunc("/user/desk/create", createDesk).Methods("POST")

	// Public
	router.HandleFunc("/users", getAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getOneUser).Methods("GET")


	log.Fatal(http.ListenAndServe(":8000", router))
}
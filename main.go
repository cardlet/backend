package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var config Config
var db *gorm.DB

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API is online")
}

func getOneUser(w http.ResponseWriter, r *http.Request) {
	var user User
	db.Find(&user, "ID = ?", mux.Vars(r)["id"])
	json.NewEncoder(w).Encode(user)
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
	user, ok := validateUser(r)
	if !ok {
		return
	}

	var updatedUser User

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the user title and description only in order to update")
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
		fmt.Fprintf(w, "Kindly enter data with the user title and description only in order to update")
	}
	
	json.Unmarshal(reqBody, &newUser)
	newUser.Token = GenerateSecureToken(20)
	db.Create(&newUser)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser.Token)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	user, ok := validateUser(r)
	if !ok {
		fmt.Fprintf(w, "Access denied")
		return
	}
	json.NewEncoder(w).Encode(user)
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
	db.Model(&user).Association("Desks").Append(desk)

	fmt.Println(desk)

	w.WriteHeader(http.StatusCreated)
}

func getConfig() Config {
	f, err := os.Open("config.yml")
	if err != nil {
    	fmt.Println("Config file not found or no permissions!")
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
    	fmt.Println("Failed to decode the config!")
	}
	return cfg
}

func main() {
	config = getConfig()

	var err error
	connection := postgres.Open("postgres://" + config.Database.Username + "@localhost/" + config.Database.Name)
	db, err = gorm.Open(connection, &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	  }
	
	db.AutoMigrate(&User{}, &Desk{}, &Card{})	
	db.Create(&User{
		Name: "Zweiter",
		Bio:  "Ne",
		Pass: "",
		Token: "1",
	})

	createRouter()
}
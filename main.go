package main

import (
	"fmt"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var config Config
var db *gorm.DB

func homeLink(w http.ResponseWriter, r *http.Request) {
	createMessageResponse(w, "API is online")
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
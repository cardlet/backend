package main

import (
	"context"
	"net/http"

	"mvdan.cc/sh/expand"
	"mvdan.cc/sh/shell"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var config map[string]expand.Variable
var db *gorm.DB

func homeLink(w http.ResponseWriter, r *http.Request) {
	createMessageResponse(w, "API is online")
}

func getDsn() string {
	dsn := "host = " + config["DB_HOST"].String() + " dbname = " + config["DB_NAME"].String() + " user = " + config["DB_USER"].String() + " password = " + config["DB_PASSWORD"].String()
	return dsn
}

func main() {
	config, _ = shell.SourceFile(context.TODO(), ".env")

	var err error
	connection := postgres.Open(getDsn())
	db, err = gorm.Open(connection, &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	  }
	
	db.AutoMigrate(&User{}, &Desk{}, &Card{})

	createRouter()
}
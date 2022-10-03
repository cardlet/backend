package api

import (
	"context"

	"github.com/cardlet/obj"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mvdan.cc/sh/expand"
	"mvdan.cc/sh/shell"
)

var config map[string]expand.Variable
var db *gorm.DB

func SetupConfig() {
	config, _ = shell.SourceFile(context.TODO(), ".env")
}

func getDsn() string {
	dsn := "host = " + config["DB_HOST"].String() + " dbname = " + config["DB_NAME"].String() + " user = " + config["DB_USER"].String() + " password = " + config["DB_PASSWORD"].String() + " port = " + config["DB_PORT"].String()
	return dsn
}

func InitDb() {
	var err error
	connection := postgres.Open(getDsn())
	db, err = gorm.Open(connection, &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	  }
	
	db.AutoMigrate(&obj.User{}, &obj.Desk{}, &obj.Card{})
}

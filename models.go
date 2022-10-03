package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name 		string 	`json:"name"`
	Bio 		string 	`json:"bio"`
	Pass		string 	`json:"password"`
	Token		string 	`json:"token"`
}

type Desk struct {
	gorm.Model
	UserID		uint   	`json:"userId"`
	Name		string 	`json:"name"`
}

type Card struct {
	gorm.Model
	UserID		uint	`json:"userId"`
	DeskID		uint	`json:"deskId"`
	Question	string 	`json:"quest"`
	Answer		string 	`json:"answer"`
}

type Config struct {
    Server struct {
        Port 	string 	`yaml:"port"`
    } `yaml:"server"`

    Database struct {
        Username	string `yaml:"user"`
        Password 	string 	`yaml:"pass"`
		Name 		string 	`yaml:"name"`
    } `yaml:"database"`
}

type Message struct {
	Message		string `json:"message"`
}

type ErrorMessage struct {
	Error		string `json:"error"`
}
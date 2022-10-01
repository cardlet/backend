package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name 	string `json:"name"`
	Bio 	string `json:"bio"`
	Pass	string `json:"password"`
	Token	string
	Desks	[]Desk `json:"desks"`
}

type Desk struct {
	gorm.Model
	Name	string `json:"name"`
	Cards	[]Card `json:"cards"`
}

type Card struct {
	gorm.Model
	Quest	string	`json:"quest"`
	Answer	string	`json:"answer"`
}
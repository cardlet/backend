package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name 	string `json:"name"`
	Bio 	string `json:"bio"`
	Pass	string `json:"password"`
	Token	string `json:"token"`
	Desks	[]Desk `json:"desks" gorm:"-"`
}

type Desk struct {
	gorm.Model
	Name	string `json:"name"`
	Cards	[]Card `json:"cards" gorm:"-"`
}

type Card struct {
	gorm.Model
	Quest	string `json:"quest"`
	Answer	string `json:"answer"`
}
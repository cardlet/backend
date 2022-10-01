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

type Config struct {
    Server struct {
        Port string `yaml:"port"`
    } `yaml:"server"`

    Database struct {
        Username string `yaml:"user"`
        Password string `yaml:"pass"`
		Name string `yaml:"name"`
    } `yaml:"database"`
}

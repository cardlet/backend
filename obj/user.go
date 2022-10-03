package obj

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name 		string 	`json:"name"`
	Bio 		string 	`json:"bio"`
	Pass		string 	`json:"password"`
	Token		string 	`json:"token"`
}
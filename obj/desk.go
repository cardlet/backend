package obj

import "gorm.io/gorm"

type Desk struct {
	gorm.Model
	UserID		uint   	`json:"userId"`
	Name		string 	`json:"name"`
}
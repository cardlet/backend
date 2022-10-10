package obj

import "gorm.io/gorm"

type Card struct {
	gorm.Model
	UserID		uint	`json:"userId"`
	DeskID		uint	`json:"deskId"`
	Question	string 	`json:"question"`
	Solution	string 	`json:"solution"`
}
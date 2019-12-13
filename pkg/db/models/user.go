package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	GoogleID	string `json:"-"`
	Email		string `gorm:"unique; not null"json:"email"`
	Name 		string `json:"name"`
	GivenName 	string `json:"given_name"`
	FamilyName 	string `json:"family_name"`
	Link 		string `json:"link"`
	Picture 	string `json:"picture"`
	Gender 		string `json:"gender"`
	Locale 		string `json:"locale"`
}

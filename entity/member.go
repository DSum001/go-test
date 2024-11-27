package entity

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	PhoneNumber string `valid:"required~phone_number is required, stringlength(10|10)~Phone Number length is not 10 digits." gorm:"uniqueIndex"`
	Password    string `valid:"required~Password is required"`
	Url         string `valid:"required~Url is required, url~Url is invalid."`
}

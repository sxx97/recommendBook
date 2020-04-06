package modles

import "github.com/jinzhu/gorm"

type UserInfo struct {
	gorm.Model
	Name string
	Email string
	Age string
	Gender string
	AccountId int16
	Password string
	Intro string
}
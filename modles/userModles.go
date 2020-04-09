package modles

import "github.com/jinzhu/gorm"

type UserInfo struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Age int `json:"age"`
	Gender string	`json:"gender"`
	AccountId int	`json:"account_id"`
	Intro string	`json:"intro"`
}

// 创建用户信息
func(user *UserInfo) CreateUser(db *gorm.DB) []error {
	errs := db.Create(user).GetErrors()
	return errs
}

// 更新用户信息
func(user *UserInfo) UpdateUserInfo(db *gorm.DB) []error {
	errs := db.Update(user).GetErrors()
	return errs
}

// 查询用户信息
func(user *UserInfo) FirstUserInfo(db *gorm.DB) []error {
	errs := db.Where("id = ?", user.Id).GetErrors()
	return errs
}
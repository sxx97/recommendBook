package modles

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Account struct {
	Id int `json:"id"`
	Account string	`json:"account"`
	Password string	`json:"password"`
	UserId int	`json:"user_id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
	DeleteTime *time.Time `json:"delete_time"`
}

func (account *Account) AddAccount(db *gorm.DB) *Account {
	nowTime := time.Now().UTC()
	createAccount := &Account{
		Account: account.Account,
		Password: account.Password,
		CreateTime: nowTime,
		UpdateTime: nowTime,
	}
	db.Create(createAccount)
	return createAccount
}

func (account Account) HasAccount(db *gorm.DB) bool {
	fmt.Println( account.Account, "查询时的account", &account)
	hasData := db.Where("account = ?", account.Account).First(&account).RecordNotFound()
	return !hasData
}

func (account Account) FirstAccount(db *gorm.DB) (result Account, err []error) {
	err = db.Where("account = ?", account.Account).First(&result).GetErrors()
	return
}
package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*var db *gorm.DB*/

/*func init() {
	db, err := gorm.Open("mysql", "root:13184234719@Mysql@(116.62.213.108:3306)/jian_shu?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接数据库错误: ", err)
	} else {
		fmt.Println("连接数据库成功")
	}
	db.SingularTable(true)
	db.DB().SetMaxOpenConns(20)
	db.DB().SetMaxOpenConns(5)
	db.LogMode(true)
}*/


func GetDB() *gorm.DB{
	db, err := gorm.Open("mysql", "root:13184234719@Mysql@(116.62.213.108:3306)/jian_shu?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接数据库错误: ", err)
	} else {
		fmt.Println("连接数据库成功")
	}
	db.SingularTable(true)
	db.DB().SetMaxOpenConns(20)
	db.DB().SetMaxOpenConns(5)
	db.LogMode(true)
	return db
}


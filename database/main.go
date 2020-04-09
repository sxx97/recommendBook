package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"os/exec"
	"path/filepath"
)

type DBConfig struct {
	DBHost string	`yaml:"DBHost"`
	DBName string	`yaml:"DBName"`
	UserName string	`yaml:"userName"`
	Password string	`yaml:"password"`
}

type SqlTypeConfig struct {
	Mysql map[string] DBConfig
}
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

func getConfigYaml() DBConfig {
	file, _ := exec.LookPath("../config/database.yaml")
	yamlPath, _ := filepath.Abs(file)
	yamlFile, err := ioutil.ReadFile(yamlPath + "/config/database.yaml")
	if err != nil {
		fmt.Println("读取数据库配置文件错误", err)
	}

	conf := new(SqlTypeConfig)
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		fmt.Println("读取数据库配置解码时错误", err)
	}
	return conf.Mysql["default"]
}


func GetDB() *gorm.DB{
	dbConfig := getConfigYaml()
	db, err := gorm.Open("mysql", "" + dbConfig.UserName + ":" + dbConfig.Password + "@(" + dbConfig.DBHost + ")/" + dbConfig.DBName + "?charset=utf8&parseTime=True&loc=Local")
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


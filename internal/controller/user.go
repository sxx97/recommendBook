package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"readBook/database"
	"readBook/modles"
	"readBook/tools"
)

const (
	AseEncryptKey string = "pas______encrypt"
)

// 注册账号处理
func RegisterHandler(ctx iris.Context) {
	accountJson := &modles.Account{}
	err := ctx.ReadJSON(accountJson)
	if err != nil {
		return
	}

	msg := "注册成功"
	statusStr := "success"

	if accountJson.Account == "" {
		msg = "账号不能为空"
		statusStr = "fail"
	}
	if accountJson.Password == "" {
		msg = "密码不能为空"
		statusStr = "fail"
	}
	// 加密操作
	encrypt := tools.AesEncrypt(accountJson.Password, AseEncryptKey)
	accountInfo := modles.Account{
		Account:  accountJson.Account,
		Password: encrypt,
	}

	db := database.GetDB()
	hasAccount := accountInfo.HasAccount(db)
	if hasAccount {
		msg = "账号已存在"
		statusStr = "fail"
	} else {
		accountInfo.AddAccount(db)
	}
	ctx.JSON(tools.ApiResource(statusStr, nil, msg))
}

// 登录处理
func LoginHandler(ctx iris.Context) {
	accountJson := &modles.Account{}

	ctx.ReadJSON(accountJson)
	msg := "登陆成功"
	statusStr := "success"
	if accountJson.Account == "" {
		msg = "账号不能为空"
		statusStr = "fail"
	}
	if accountJson.Password == "" {
		msg = "密码不能为空"
		statusStr = "fail"
	}
	accountJson.Password = tools.AesEncrypt(accountJson.Password, AseEncryptKey)
	fmt.Println(accountJson, "accountJson的内容")
	db := database.GetDB()
	findAccount, errs := accountJson.FirstAccount(db)
	if len(errs) > 0 {
		msg = errs[0].Error()
		fmt.Println("查看查询失败的原因", errs)
		statusStr = "fail"
	} else {
		if findAccount.Password != accountJson.Password {
			msg = "账号或密码错误"
			statusStr = "fail"
		}
	}
	token := tools.CreateJWTToken(findAccount.Id, findAccount.Account)
	if statusStr == "fail" {
		token = ""
	}
	ctx.JSON(tools.ApiResource(statusStr, token, msg))
}

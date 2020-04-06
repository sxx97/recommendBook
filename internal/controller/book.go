package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"readBook/database"
	"readBook/modles"
	"readBook/tools"
)

func FindBooksHandler(ctx iris.Context) {
	var (
		bookInfo modles.Book
		statusStr = "success"
		msg string
	)
	keyword := ctx.Params().Get("keyword")
	fmt.Println("keyword获取到的内容", keyword)
	db := database.GetDB()
	books, errs := bookInfo.FindBlurryBooks(db, keyword)
	if len(errs) > 0 {
		fmt.Println("搜索图书错误", errs[0].Error())
		books = nil
		statusStr = "fail"
		msg = errs[0].Error()
	}
	ctx.JSON(tools.ApiResource(statusStr, books, msg))
}
package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"readBook/database"
	"readBook/modles"
	"readBook/tools"
)

// 处理图书搜索
func FindBooksHandler(ctx iris.Context) {
	var (
		bookInfo modles.Book
		statusStr = "success"
		msg string
	)
	keyword := ctx.URLParam("keyword")
	db := database.GetDB()
	books, errs := bookInfo.FindBlurryBooks(db, keyword)
	if len(errs) > 0 {
		fmt.Println("搜索图书错误", errs[0].Error())
		books = nil
		statusStr = "fail"
		msg = errs[0].Error()
	}
	var responseData map[string]interface{}
	if books != nil {
		responseData = map[string]interface{}{
			"data": books,
			"total": len(books),
		}
	}
	ctx.JSON(tools.ApiResource(statusStr, responseData, msg))
}
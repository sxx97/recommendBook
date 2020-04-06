package routes

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"readBook/internal/controller"
	"readBook/tools"
)

func Route(app *iris.Application) {
	staticPagePath(app)
	apiRequestPath(app)
}

func apiBeforeHandler (ctx iris.Context) {
	if false {
		ctx.WriteString("Authentication failed")
	} else {
		ctx.Next()
	}
}

func responseHandler(ctx iris.Context) {
	fmt.Println("处理中")
	ctx.Next()
}


func RouteError(app *iris.Application) {
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.Writef("404页面 ")
	})
}

// api开头的请求路径
func apiRequestPath(app *iris.Application) {
	api := apiPartyAndAuthorization(app)
	api.Post("/register", controller.RegisterHandler)
	api.Post("/login", controller.LoginHandler)
	api.Get("/search/book", controller.FindBooksHandler)
}

// 创建api开头的请求路径并验证授权
func apiPartyAndAuthorization(app *iris.Application) (apiParty iris.Party){
	apiParty = app.Party("/api", tools.JwtMiddleWare().Serve, responseHandler)
	return
}

// 静态页面的路径设置
func staticPagePath(app *iris.Application) {
	app.Get("/", func(ctx iris.Context){
		ctx.HTML("<p>Hello World</p>")
	})
}
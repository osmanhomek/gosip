package main

import (
	"gosip/controllers"
	"gosip/middlewares"

	"github.com/iris-contrib/middleware/logger"
	"github.com/kataras/go-template/html"
	"github.com/kataras/iris"
)

func main() {
	//CONFIG
	iris.New(
		iris.OptionCharset("UTF-8"), iris.OptionIsDevelopment(true),
		iris.OptionSessionsCookie("gosip"), iris.OptionWebsocketEndpoint("/my_endpoint"),
	)
	iris.Config.EnablePathEscape = true

	//STATIC TEMPLATE
	iris.StaticWeb("/plugins", "./static/plugins")
	iris.StaticWeb("/dist", "./static/dist")
	iris.StaticWeb("/bootstrap", "./static/bootstrap")
	iris.UseTemplate(html.New(html.Config{Layout: "layout.html"})).Directory("./templates", ".html")

	// MIDDLEWARE
	iris.UseFunc(middlewares.RequestCycleStart)
	iris.DoneFunc(middlewares.RequestCycleEnd)

	// ROUTER
	iris.Get("/", controllers.LoginCheck)
	iris.Get("/login", controllers.Login)
	iris.Get("/dashboard", controllers.Dashboard)
	iris.Get("/gittigidiyor", controllers.Gittigidiyor)

	// ERROR
	errorLogger := logger.New()
	iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
		errorLogger.Serve(ctx)
		ctx.HTML(iris.StatusNotFound, "<h1>404 hata sayfasi</h1>")
	})
	iris.OnError(iris.StatusForbidden, func(ctx *iris.Context) {
		errorLogger.Serve(ctx)
		ctx.HTML(iris.StatusForbidden, "<h1> You are not allowed here </h1>")
	})

	//RUN
	iris.Listen(":8080")
}

package controllers

import (
	"fmt"

	"github.com/kataras/iris"
)

// Dashboard : describe what this function does
func Dashboard(ctx *iris.Context) {
	fmt.Println("Action")

	ctx.Render("dashboard.html", loginvar{"My Page title", "Hello world!"}, iris.RenderOptions{"gzip": true})
}

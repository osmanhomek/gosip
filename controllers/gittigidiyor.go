package controllers

import (
	"fmt"

	"github.com/kataras/iris"
)

// Gittigidiyor : describe what this function does
func Gittigidiyor(ctx *iris.Context) {
	fmt.Println("Action")

	ctx.Render("gittigidiyor.html", loginvar{"My Page title", "Hello world!"}, iris.RenderOptions{"gzip": true})
}

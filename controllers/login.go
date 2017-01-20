package controllers

import (
	"fmt"

	"github.com/kataras/iris"
)

type loginvar struct {
	title   string
	message string
}

// Login : describe what this function does
func Login(ctx *iris.Context) {
	fmt.Println("Action")

	ctx.Render("login.html", loginvar{"My Page title", "Hello world!"}, iris.RenderOptions{"gzip": true})
}

// LoginCheck :
func LoginCheck(ctx *iris.Context) {
	fmt.Println("LoginCheck")
}

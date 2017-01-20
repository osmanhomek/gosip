package middlewares

import (
	"fmt"

	"github.com/kataras/iris"
)

// SessionCheck :
func SessionCheck(c *iris.Context) {
	fmt.Println("Session")
	c.Next()
}

package middlewares

import (
	"fmt"

	"github.com/kataras/iris"
)

// RequestCycle :
func RequestCycleStart(c *iris.Context) {
	Logging(c)
	SessionCheck(c)
}

// RequestCycleEnd :
func RequestCycleEnd(ctx *iris.Context) {
	fmt.Println("RequestCycleEnd")
}

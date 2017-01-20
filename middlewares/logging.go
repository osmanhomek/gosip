package middlewares

import (
	"fmt"

	"github.com/iris-contrib/middleware/logger"
	"github.com/kataras/iris"
)

// Logging :
func Logging(c *iris.Context) {
	fmt.Println("Logging")
	sipoLogger := logger.New(logger.Config{
		Status: true,
		IP:     true,
		Method: true,
		Path:   true,
	})
	sipoLogger.Serve(c)
}

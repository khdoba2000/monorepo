package main

import (
	"fmt"
	"net/http"

	"monorepo/src/libs/hello"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/auth-service/hello", func(c *gin.Context) {
		c.String(http.StatusOK, hello.Greet("World"))
	})
	_ = r.Run(":8084")
	fmt.Println("request recieved to service one")
}

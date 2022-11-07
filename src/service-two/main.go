package main

import (
	"fmt"
	"net/http"

	"monorepo/src/libs/hello"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/two/hello", func(c *gin.Context) {
		c.String(http.StatusOK, hello.Greet("World"))
	})
	_ = r.Run(":8082")
	fmt.Println("Request recieved to service two")
}

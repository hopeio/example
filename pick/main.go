package main

import (
	"github.com/gin-gonic/gin"
	pickgin "github.com/hopeio/pick/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	pickgin.Register(server, true, &UserService{})
	server.Run(":8080")
}

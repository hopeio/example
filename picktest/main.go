package main

import (
	"github.com/gin-gonic/gin"
	pickgin "github.com/hopeio/pick/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	pickgin.RegisterService(&UserService{})
	server := gin.New()
	pickgin.Start(server, true, "test", false)
	server.Run(":8080")
}

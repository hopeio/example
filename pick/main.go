package main

import (
	"github.com/gin-gonic/gin"
	pickgin "github.com/hopeio/pick/gin"
	"log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	pickgin.Register(server, true, &UserService{})
	log.Fatal(server.Run(":8080"))
}

package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Log(ctx *gin.Context) {
	log.Println(ctx.Request.Method, ctx.Request.RequestURI)
}

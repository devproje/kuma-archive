package middleware

import "github.com/gin-gonic/gin"

func Header(ctx *gin.Context) {
	ctx.Header("Server", "Golang,Gin")
	ctx.Next()
}

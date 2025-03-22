package middleware

import "github.com/gin-gonic/gin"

func Header(ctx *gin.Context) {
	ctx.Header("X-Powered-By", "Golang")
	ctx.Next()
}

package middleware

import (
	"fmt"
	"os"
	"strings"

	"git.wh64.net/devproje/kuma-archive/internal/service"
	"github.com/gin-gonic/gin"
)

func BasicAuth(ctx *gin.Context) {
	var matches = false
	var list = []string{"/settings"}

	for _, i := range list {
		if !strings.Contains(ctx.Request.URL.Path, i) {
			continue
		}

		matches = true
	}

	if !matches {
		ctx.Next()
		return
	}

	auth := service.NewAuthService()
	username, password, ok := ctx.Request.BasicAuth()
	if !ok {
		ctx.Status(403)
		return
	}

	ok, err := auth.VerifyToken(username, password)
	if err != nil {
		ctx.Status(500)
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}

	if !ok {
		ctx.Status(403)
		return
	}

	ctx.Next()
}

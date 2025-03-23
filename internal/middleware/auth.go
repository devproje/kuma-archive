package middleware

import (
	"fmt"
	"os"
	"strings"

	"git.wh64.net/devproje/kuma-archive/internal/service"
	"github.com/gin-gonic/gin"
)

func WorkerRoute(ctx *gin.Context) {
	if !strings.HasPrefix(ctx.Request.URL.Path, "/api/worker") {
		ctx.Next()
		return
	}

	var err error
	var dirs []service.PrivDir
	auth := service.NewAuthService()
	privdir := service.NewPrivDirService(nil)
	dirs = privdir.Query()
	if len(dirs) == 0 {
		ctx.Next()
		return
	}

	var target string
	var matches = false
	for _, dir := range dirs {
		if !strings.HasSuffix(ctx.Request.URL.Path, dir.DirName) {
			continue
		}

		target = dir.DirName
		matches = true
	}

	if !matches {
		ctx.Next()
		return
	}

	username, password, ok := ctx.Request.BasicAuth()
	if !ok {
		ctx.JSON(401, gin.H{
			"ok":    0,
			"errno": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	ok, err = auth.VerifyToken(username, password)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		ctx.JSON(401, gin.H{
			"ok":    0,
			"errno": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	var acc *service.Account
	acc, err = auth.Read(username)
	if err != nil {
		ctx.JSON(500, gin.H{
			"ok":    0,
			"errno": "Internal Server Error",
		})

		ctx.Abort()
		return
	}

	privdir = service.NewPrivDirService(acc)
	if !ok {
		ctx.JSON(401, gin.H{
			"ok":    0,
			"errno": "Unauthorized",
		})

		ctx.Abort()
		return
	}

	var d *service.PrivDir
	d, err = privdir.Read(target)
	if err != nil {
		ctx.JSON(500, gin.H{
			"ok":    0,
			"errno": "Internal Server Error",
		})
		ctx.Abort()
		return
	}

	if d == nil {
		ctx.JSON(401, gin.H{
			"ok":    0,
			"errno": "Unauthorized",
		})
		ctx.Abort()
		return
	}

	ctx.Next()
}

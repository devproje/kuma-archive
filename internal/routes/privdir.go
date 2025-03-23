package routes

import (
	"fmt"
	"git.wh64.net/devproje/kuma-archive/internal/service"
	"github.com/gin-gonic/gin"
	"os"
)

func createDir(ctx *gin.Context) {
	var err error
	auth := service.NewAuthService()
	username, password, ok := ctx.Request.BasicAuth()
	if !ok {
		ctx.JSON(401, gin.H{
			"ok":    0,
			"errno": "Unauthorized",
		})
		return
	}

	if ok, err = auth.VerifyToken(username, password); !ok {
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}

		ctx.JSON(401, gin.H{
			"ok":    0,
			"errno": "Unauthorized",
		})
		return
	}

	var acc *service.Account
	acc, err = auth.Read(username)
	if err != nil {
		ctx.JSON(500, gin.H{
			"ok":    0,
			"errno": "Interval Server Error",
		})
		return
	}

	path := ctx.PostForm("path")
	privdir := service.NewPrivDirService(acc)

	var id string
	if id, err = privdir.Create(path); err != nil {
		ctx.JSON(500, gin.H{
			"ok":    0,
			"errno": fmt.Sprintf("'%s' directory is already registered", path),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"ok":     1,
		"dir_id": id,
	})
}

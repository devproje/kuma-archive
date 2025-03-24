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
		ctx.Status(401)
		return
	}

	if ok, err = auth.VerifyToken(username, password); !ok {
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}

		ctx.Status(401)
		return
	}

	var acc *service.Account
	acc, err = auth.Read(username)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		ctx.Status(500)
		return
	}

	path := ctx.PostForm("path")
	privdir := service.NewPrivDirService(acc)

	var id string
	if id, err = privdir.Create(path); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		ctx.Status(500)
		return
	}

	ctx.JSON(200, gin.H{
		"ok":     1,
		"dir_id": id,
	})
}

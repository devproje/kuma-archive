package routes

import (
	"fmt"
	"os"

	"git.wh64.net/devproje/kuma-archive/internal/service"
	"github.com/gin-gonic/gin"
)

func readAcc(ctx *gin.Context) {
	auth := service.NewAuthService()
	username, password, ok := ctx.Request.BasicAuth()
	if !ok {
		ctx.Status(401)
		return
	}

	ok, err := auth.VerifyToken(username, password)
	if err != nil {
		ctx.JSON(500, gin.H{
			"ok":    0,
			"errno": "internal server error!",
		})
	}

	if !ok {
		ctx.JSON(401, gin.H{
			"ok":    0,
			"errno": "unauthorized",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"ok":       1,
		"username": username,
	})
}

func updateAcc(ctx *gin.Context) {
	auth := service.NewAuthService()
	old := ctx.PostForm("password")
	pass := ctx.PostForm("new_password")
	username, _, ok := ctx.Request.BasicAuth()
	if !ok {
		ctx.Status(401)
		return
	}

	ok, err := auth.Verify(username, old)
	if err != nil || !ok {
		ctx.Status(401)
		return
	}

	if err = auth.Update(username, pass); err != nil {
		ctx.Status(500)
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}

	ctx.Status(200)
}

func deleteAcc(ctx *gin.Context) {
	auth := service.NewAuthService()
	username, password, ok := ctx.Request.BasicAuth()
	if !ok {
		ctx.Status(401)
		return
	}

	ok, err := auth.VerifyToken(username, password)
	if err != nil {
		ctx.Status(500)
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}

	if !ok {
		ctx.Status(401)
		return
	}

	if err = auth.Delete(username); err != nil {
		ctx.Status(500)
		_, _ = fmt.Fprintln(os.Stderr, err)
		return
	}

	ctx.Status(200)
}

func login(ctx *gin.Context) {
	auth := service.NewAuthService()
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	acc, err := auth.Read(username)
	if err != nil {
		ctx.JSON(401, gin.H{
			"ok":    0,
			"errno": "username or password not invalid",
		})
		return
	}

	ok, err := auth.Verify(username, password)
	if err != nil || !ok {
		ctx.JSON(401, gin.H{
			"ok":    0,
			"errno": "username or password not invalid",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"ok":    1,
		"token": auth.Token(acc.Username, acc.Password),
	})
}

func check(ctx *gin.Context) {
	auth := service.NewAuthService()
	username, password, ok := ctx.Request.BasicAuth()
	if !ok {
		ctx.Status(401)
		return
	}

	ok, err := auth.VerifyToken(username, password)
	if err != nil || !ok {
		ctx.Status(401)
		return
	}

	ctx.Status(200)
}

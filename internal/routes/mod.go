package routes

import (
	"git.wh64.net/devproje/kuma-archive/internal/middleware"
	"git.wh64.net/devproje/kuma-archive/internal/service"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func New(app *gin.Engine, version *service.Version, apiOnly bool) {
	app.Use(middleware.CORS)
	app.Use(middleware.Header)
	app.Use(middleware.BasicAuth)

	api := app.Group("/api")
	api.GET("/path/*path", discoverPath)
	api.GET("/download/*path", downloadPath)

	auth := api.Group("/auth")
	{
		auth.POST("/login", login)
		auth.GET("/read", readAcc)
		auth.PATCH("/update", updateAcc)
		auth.DELETE("/delete", deleteAcc)
	}

	api.GET("/version", func(ctx *gin.Context) {
		ctx.String(200, "%s", version.String())
	})

	if apiOnly {
		return
	}

	app.Use(static.Serve("/", static.LocalFile("./web", true)))
	app.Use(static.Serve("/assets", static.LocalFile("./assets", false)))

	app.NoRoute(func(ctx *gin.Context) {
		ctx.File("./web/index.html")
	})

	app.GET("favicon.ico", func(ctx *gin.Context) {
		ctx.File("/web/assets/favicon.ico")
	})
}

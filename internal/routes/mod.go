package routes

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func New(app *gin.Engine) {
	app.Use(static.Serve("/", static.LocalFile("./public", true)))
	app.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", nil)
	})

	api := app.Group("/api")
	{
		api.GET("/path/*path", func(ctx *gin.Context) {
			path := ctx.Param("path")

			ctx.JSON(200, gin.H{
				"ok":   1,
				"path": path,
			})
		})
	}
}

package routes

import (
	"fmt"
	"os"

	"git.wh64.net/devproje/kuma-archive/internal/service"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func New(app *gin.Engine, apiOnly bool) {
	api := app.Group("/api")
	{
		api.GET("/path/*path", func(ctx *gin.Context) {
			worker := service.NewWorkerService()

			path := ctx.Param("path")
			data, err := worker.Read(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				ctx.Status(404)
				return
			}

			if !data.IsDir {
				ctx.FileAttachment(data.Path, data.Name)
				return
			}

			raw, err := os.ReadDir(data.Path)
			if err != nil {
				ctx.Status(500)
				return
			}

			entries := make([]service.DirEntry, 0)
			for _, entry := range raw {
				finfo, err := entry.Info()
				if err != nil {
					fmt.Fprintf(os.Stderr, "%v\n", err)
					continue
				}

				entries = append(entries, service.DirEntry{
					Name:     entry.Name(),
					Path:     path,
					FileSize: uint64(finfo.Size()),
					IsDir:    finfo.IsDir(),
				})
			}

			ctx.JSON(200, gin.H{
				"ok":      1,
				"path":    path,
				"entries": entries,
			})
		})
	}

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

package routes

import (
	"fmt"
	"os"

	"git.wh64.net/devproje/kuma-archive/internal/service"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func New(app *gin.Engine) {
	app.Use(static.Serve("/", static.LocalFile("./public", true)))
	app.Use(static.Serve("/assets", static.LocalFile("./assets", false)))

	app.NoRoute(func(ctx *gin.Context) {
		ctx.File("./public/index.html")
	})

	app.GET("favicon.ico", func(ctx *gin.Context) {
		ctx.File("/assets/favicon.ico")
	})

	api := app.Group("/api")
	{
		api.GET("/path/*path", func(ctx *gin.Context) {
			worker := service.NewWorkerService()

			path := ctx.Param("path")
			info, err := worker.Read(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				ctx.Status(404)
				return
			}

			if !info.IsDir {
				ctx.FileAttachment(info.FullPath, info.Name)
				return
			}

			raw, err := os.ReadDir(info.FullPath)
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
					FileSize: uint64(finfo.Size()),
					FullPath: path,
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
}

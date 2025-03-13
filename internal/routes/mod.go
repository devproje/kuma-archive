package routes

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"git.wh64.net/devproje/kuma-archive/config"
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

			if !info.IsDir() {
				var split = strings.Split(path, "/")
				ctx.FileAttachment(filepath.Join(config.INDEX_DIR, path), split[len(split)-1])

				return
			}

			entry, err := os.ReadDir(filepath.Join(config.INDEX_DIR, path))
			if err != nil {
				ctx.Status(500)
				return
			}

			entries := make([]service.DirEntry, 0)
			for _, fd := range entry {
				var info, _ = fd.Info()
				entries = append(entries, service.DirEntry{
					Name:     fd.Name(),
					FileSize: uint64(info.Size()),
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

package routes

import (
	"fmt"
	"os"
	"path/filepath"

	"git.wh64.net/devproje/kuma-archive/internal/middleware"
	"git.wh64.net/devproje/kuma-archive/internal/service"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func New(app *gin.Engine, version *service.Version, apiOnly bool) {
	app.Use(middleware.CORS)
	app.Use(middleware.BasicAuth)

	api := app.Group("/api")
	api.GET("/path/*path", readPath)
	api.GET("/download/*path", downloadPath)

	api.POST("/private")

	authentication(api.Group("/auth"))

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

func readPath(ctx *gin.Context) {
	worker := service.NewWorkerService()
	path := ctx.Param("path")
	// TODO: prefix detect
	// if strings.HasPrefix(path, "")

	data, err := worker.Read(path)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		ctx.Status(404)
		return
	}

	if !data.IsDir {
		ctx.JSON(200, gin.H{
			"ok":      1,
			"path":    path,
			"total":   data.FileSize,
			"is_dir":  false,
			"entries": nil,
		})
		return
	}

	raw, err := os.ReadDir(data.Path)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
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
			Path:     filepath.Join(path, entry.Name()),
			Date:     finfo.ModTime().Unix(),
			FileSize: uint64(finfo.Size()),
			IsDir:    finfo.IsDir(),
		})
	}

	ctx.JSON(200, gin.H{
		"ok":      1,
		"path":    path,
		"total":   data.FileSize,
		"is_dir":  true,
		"entries": entries,
	})
}

func downloadPath(ctx *gin.Context) {
	worker := service.NewWorkerService()
	path := ctx.Param("path")
	data, err := worker.Read(path)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		ctx.Status(404)
		return
	}

	if data.IsDir {
		ctx.String(400, "current path is not file")
		return
	}

	ctx.FileAttachment(data.Path, data.Name)
}

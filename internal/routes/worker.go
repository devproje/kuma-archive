package routes

import (
	"fmt"
	"git.wh64.net/devproje/kuma-archive/internal/service"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

func readPath(ctx *gin.Context) {
	worker := service.NewWorkerService()
	path := ctx.Param("path")

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
		var finfo os.FileInfo
		finfo, err = entry.Info()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
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

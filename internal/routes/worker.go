package routes

import (
	"fmt"
	"git.wh64.net/devproje/kuma-archive/internal/service"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

func worker(path string) (*service.DirEntry, error) {
	sv := service.NewWorkerService()
	return sv.Read(path)
}

func discoverPath(ctx *gin.Context) {
	path := ctx.Param("path")
	data, err := worker(path)
	if data == nil {
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}

		ctx.JSON(404, gin.H{
			"ok":    0,
			"errno": fmt.Errorf("path %s is not exist", path),
		})
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

	var raw []os.DirEntry
	raw, err = os.ReadDir(data.Path)
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
	path := ctx.Param("path")
	data, err := worker(path)
	if data == nil {
		if err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}

		ctx.JSON(404, gin.H{
			"ok":    0,
			"errno": fmt.Errorf("path %s is not exist", path),
		})
		return
	}

	if data.IsDir {
		ctx.JSON(404, gin.H{
			"ok":    0,
			"errno": "file is not exist",
		})
		return
	}

	ctx.FileAttachment(data.Path, data.Name)
}

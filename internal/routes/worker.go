package routes

import (
	"fmt"
	"git.wh64.net/devproje/kuma-archive/internal/service"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strings"
)

func checkAuth(ctx *gin.Context) (bool, error) {
	privdir := service.NewPrivDirService(nil)
	dirs, err := privdir.Query()
	if err != nil {
		return true, nil
	}

	for _, dir := range dirs {
		if !strings.HasPrefix(ctx.Request.URL.Path, dir.DirName) {
			continue
		}

		auth := service.NewAuthService()
		username, password, ok := ctx.Request.BasicAuth()
		if !ok {
			return false, nil
		}

		ok, err = auth.VerifyToken(username, password)
		if err != nil {
			return false, err
		}

		if !ok {
			return false, nil
		}

		var acc *service.Account
		acc, err = auth.Read(username)
		if err != nil {
			return false, err
		}

		var path *service.PrivDir
		privdir = service.NewPrivDirService(acc)
		path, err = privdir.Read(dir.DirName)
		if err != nil {
			return false, err
		}

		if path == dir {
			return true, nil
		}

		return false, nil
	}

	return true, nil
}

func readPath(ctx *gin.Context) {
	ok, err := checkAuth(ctx)
	if err != nil {
		ctx.Status(401)
		return
	}

	if !ok {
		ctx.Status(401)
		return
	}

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
	ok, err := checkAuth(ctx)
	if err != nil {
		ctx.Status(401)
		return
	}

	if !ok {
		ctx.Status(401)
		return
	}

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

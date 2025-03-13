package service

import (
	"os"
	"path/filepath"

	"git.wh64.net/devproje/kuma-archive/config"
)

type WorkerService struct{}

type DirEntry struct {
	Name string `json:"name"`
}

func NewWorkerService() *WorkerService {
	return &WorkerService{}
}

func (sv *WorkerService) Read(path string) (os.FileInfo, error) {
	fullpath := filepath.Join(config.INDEX_DIR, path)
	return os.Stat(fullpath)
}

package service

import (
	"os"
	"path/filepath"

	"git.wh64.net/devproje/kuma-archive/config"
)

type WorkerService struct{}

type DirEntry struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Date     int64  `json:"date"`
	FileSize uint64 `json:"file_size"`
	IsDir    bool   `json:"is_dir"`
}

func NewWorkerService() *WorkerService {
	return &WorkerService{}
}

func (sv *WorkerService) Read(path string) (*DirEntry, error) {
	fullpath := filepath.Join(config.INDEX_DIR, path)
	info, err := os.Stat(fullpath)
	if err != nil {
		return nil, err
	}

	ret := DirEntry{
		Name:     info.Name(),
		Date:     info.ModTime().Unix(),
		Path:     fullpath,
		FileSize: uint64(info.Size()),
		IsDir:    info.IsDir(),
	}
	return &ret, nil
}

package service

import (
	"os"
	"path/filepath"

	"git.wh64.net/devproje/kuma-archive/config"
)

type WorkerService struct{}

type DirEntry struct {
	Name     string `json:"name"`
	FullPath string `json:"path"`
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
		FullPath: fullpath,
		FileSize: uint64(info.Size()),
	}
	return &ret, nil
}

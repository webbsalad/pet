package file_storage

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/webbsalad/pet/internal/config"
	"go.uber.org/fx"
)

type FileStorage struct {
	Path string
	mu   sync.Mutex
}

func InitFileStorage(cfg config.Config, lc fx.Lifecycle) (*FileStorage, error) {
	dir := filepath.Dir(cfg.STORAGE_PATH)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	f, err := os.OpenFile(cfg.STORAGE_PATH, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	f.Close()

	fs := &FileStorage{Path: cfg.STORAGE_PATH}
	log.Printf("File storage configured: %s", cfg.STORAGE_PATH)

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			log.Println("closing FileStorage")
			return nil
		},
	})

	return fs, nil
}

func (fs *FileStorage) Set(v interface{}) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	f, err := os.OpenFile(fs.Path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}

func (fs *FileStorage) Get(v interface{}) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	f, err := os.OpenFile(fs.Path, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewDecoder(f).Decode(v)
}

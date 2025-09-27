package data

import (
	"curaitor/internal/model"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"sync"
)

type CachedFiles struct {
	Mu         *sync.Mutex
	CachedFiles []model.CachedFile
}

func LoadCache() (*CachedFiles, error) {
	caches, err := loadCache()
	if err != nil {
		return nil, fmt.Errorf("failed to load cache: %w", err)
	}

	slog.Info("cache loaded")

	return &CachedFiles{
		Mu:         &sync.Mutex{},
		CachedFiles: caches,
	}, nil
}

func loadCache() ([]model.CachedFile, error) {
	var caches []model.CachedFile 

	data, err := os.ReadFile("cache.json")
	if err != nil {
		if os.IsNotExist(err) {
			return caches, nil
		}
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	if len(data) == 0 {
		return caches, nil
	}

	if err := json.Unmarshal(data, &caches); err != nil {
		return nil, fmt.Errorf("failed to unmarshal cache: %w", err)
	}

	return caches, nil
}

func (c *CachedFiles) Add(cachedFile model.CachedFile) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.CachedFiles = append(c.CachedFiles, cachedFile)
	slog.Info("file cached", slog.String("file", cachedFile.FilePath))
}

func (c *CachedFiles) Save() error {
	c.Mu.Lock() 
	defer c.Mu.Unlock() 
	dataBytes, err := json.MarshalIndent(c.CachedFiles, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal cache: %w", err)
	}

	if err := os.WriteFile("cache.json", dataBytes, 0644); err != nil {
		return fmt.Errorf("failed to write cache file: %w", err)
	}

	return nil
}

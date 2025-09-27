package gemini

import (
	"context"
	"curaitor/internal/config"
	"curaitor/internal/data"
	"curaitor/internal/model"
	"fmt"
	"log/slog"
	"os"
	"sync"
)

func ParseMainFileWorker(cfg *config.Config, ctx context.Context, wg *sync.WaitGroup, caches *data.CachedFiles, newMainFilesCh <-chan string, errCh chan<- error, fileEdgeCh chan<- string) {
	defer wg.Done()
	for {
		select {
		case file := <-newMainFilesCh:
			c, err := os.ReadFile(file)
			if err != nil {
				errCh <- err
				continue
			}

			if len(c) == 0 {
				slog.Warn("empty file, skipping", slog.String("file", file))
				continue
			}

			caches.Add(model.CachedFile{
				FilePath: file,
				Content:  c,
			})
			if err := caches.Save(); err != nil {
				errCh <- fmt.Errorf("failed to save cache: %w", err)
			}

			fileEdgeCh <- file

		case <-ctx.Done():
			slog.Info("ParseMainFileWorker done")
			return
		}
	}
}

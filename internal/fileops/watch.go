package fileops

import (
	"context"
	"curaitor/internal/config"
	"fmt"
	"log/slog"
	"time"

	"github.com/radovskyb/watcher"
)

func StartDumpWatcher(cfg *config.Config, ctx context.Context, newFilesCh chan<- string, errCh chan<- error) {
	w := watcher.New()
	w.FilterOps(watcher.Create)

	if err := w.AddRecursive(cfg.WatcherPath); err != nil {
		errCh <- fmt.Errorf("failed to add watcher path: %w", err)
	}

	go watcherLoop(ctx, newFilesCh, w)

	slog.Info("watcher started")

	if err := w.Start(time.Duration(cfg.WatcherIntervalSeconds) * time.Second); err != nil {
		errCh <- fmt.Errorf("failed to start watcher: %w", err)
	}
}


func watcherLoop(ctx context.Context, newFilesCh chan<- string, w *watcher.Watcher) {
	w.Wait()

	for {
		select {
		case event := <-w.Event:
			if !event.IsDir() {
				slog.Info("file added", slog.String("file", event.Path))
				newFilesCh <- event.Path
			}
		case err := <-w.Error:
			slog.Error("error in watcher", slog.Any("error", err))
		case <-ctx.Done():
			slog.Info("watcher done")
			w.Close()
			return
		}
	}
}

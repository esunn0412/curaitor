package fileops

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/radovskyb/watcher"
)

func StartWatcher(path string, watcherIntervalSeconds int, ctx context.Context, newMainFilesCh chan<- string, errCh chan<- error) {
	w := watcher.New() 
	w.FilterOps(watcher.Create) 

	if err := w.AddRecursive(path); err != nil {
		errCh <- fmt.Errorf("failed to add watcher path: %w", err)
	}

	go watcherLoop(ctx, newMainFilesCh, w)

	slog.Info("main watcher started")

	if err := w.Start(time.Duration(watcherIntervalSeconds) * time.Second); err != nil {
		errCh <- fmt.Errorf("failed to start watcher: %w", err)
	}
}

func watcherLoop(ctx context.Context, newDumpFilesCh chan<- string, w *watcher.Watcher) {
	w.Wait()

	for {
		select {
		case event := <-w.Event:
			if !event.IsDir() {
				slog.Info("file added", slog.String("file", event.Path))
				newDumpFilesCh <- event.Path
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

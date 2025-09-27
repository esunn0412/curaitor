package main

import (
	"context"
	"curaitor/internal/config"
	"curaitor/internal/data"
	"curaitor/internal/fileops"
	"curaitor/internal/gemini"

	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		slog.Error("failed to load config", slog.Any("error", err))
		os.Exit(1)
	}

	courses, err := data.LoadCourses()
	if err != nil {
		slog.Error("failed to load courses", slog.Any("error", err))
		os.Exit(1)
	}

	var (
		newFilesCh  = make(chan string)
		errCh       = make(chan error)
		wg          = &sync.WaitGroup{}
		ctx, cancel = signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	)
	defer cancel()

	go fileops.StartDumpWatcher(cfg, ctx, newFilesCh, errCh)

	for range cfg.NumParseFileWorkers {
		wg.Add(1)
		go gemini.ParseFileWorker(cfg, ctx, wg, courses, newFilesCh, errCh)
	}

	heartbeat := time.NewTicker(time.Duration(cfg.HeartbeatIntervalSeconds) * time.Second)
	defer heartbeat.Stop()

	for {
		select {
		case <-heartbeat.C:
			slog.Info("service running (heartbeat)")
		case err := <-errCh:
			slog.Error("error", slog.Any("error", err))
		case <-ctx.Done():
			slog.Info("shutting down")
			close(newFilesCh)
			close(errCh)
			wg.Wait()
			return
		}
	}
}

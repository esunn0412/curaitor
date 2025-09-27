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

	quizzes, err := data.LoadQuiz()
	if err != nil {
		slog.Error("failed to load quizzes", slog.Any("error", err)) // Erroring
		os.Exit(1)
	}
	
	// quizzes, err := data.LoadQuiz()
	// if err != nil {
	// 	slog.Error("failed to load quizzes", slog.Any("error", err)) // Erroring
	// 	os.Exit(1)
	// }

	var (
		newDumpFilesCh = make(chan string)
		newMainFilesCh = make(chan string)
		newQuizCodesCh = make(chan string)
		errCh          = make(chan error)
		wg             = &sync.WaitGroup{}
		ctx, cancel    = signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	)
	defer cancel()

	go fileops.StartWatcher(cfg.DumpWatcherPath, cfg.WatcherIntervalSeconds, ctx, newDumpFilesCh, errCh)
	go fileops.StartWatcher(cfg.SchoolPath, cfg.WatcherIntervalSeconds, ctx, newMainFilesCh, errCh)

	for range cfg.NumParseFileWorkers {
		wg.Add(1)
		go gemini.ParseFileWorker(cfg, ctx, wg, courses, newDumpFilesCh, errCh)
	}

	for range cfg.NumGenerateQuizWorkers {
		wg.Add(1)
		go gemini.GenerateQuizWorker(cfg, ctx, wg, quizzes, newQuizCodesCh, errCh)
	}

	// Dummy get request that gives a course code
	time.Sleep(20 * time.Second)
	const code = "CS370"
	newQuizCodesCh <- code

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
			close(newDumpFilesCh)
			close(newMainFilesCh)
			// close(newQuizCodesCh)
			close(errCh)
			wg.Wait()
			return
		}
	}
}

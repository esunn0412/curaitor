package main

import (
	"context"
	"curaitor/internal/api"
	"curaitor/internal/config"
	"curaitor/internal/data"
	"curaitor/internal/fileops"
	"curaitor/internal/gemini"
	"net/http"

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

	caches, err := data.LoadCache()
	if err != nil {
		slog.Error("failed to load cache", slog.Any("error", err))
		os.Exit(1)
	}

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
		wg.Add(2)
		go gemini.ParseDumpFileWorker(cfg, ctx, wg, courses, newDumpFilesCh, errCh)
		go gemini.ParseMainFileWorker(cfg, ctx, wg, caches, newMainFilesCh, errCh)
	}

	for range cfg.NumGenerateQuizWorkers {
		wg.Add(1)
		go gemini.GenerateQuizWorker(cfg, ctx, wg, quizzes, newQuizCodesCh, errCh)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/courses", api.GetCoursesHandler(courses))
	mux.HandleFunc("/quiz", api.GetQuizHandler(quizzes, newQuizCodesCh))
	mux.HandleFunc("/quiz/regenerate", api.RegenerateQuizHandler(newQuizCodesCh))
	mux.HandleFunc("/files", api.GetFilesHandler(caches))

	server := &http.Server{
		Addr:    cfg.ServerAddr,
		Handler: mux,
	}

	go func() {
		slog.Info("server starting", slog.String("addr", cfg.ServerAddr))
		if err := server.ListenAndServe(); err != nil {
			slog.Error("failed to listen and serve", slog.Any("error", err))
			os.Exit(1)
		}
	}()

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
			shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer shutdownCancel()
			if err := server.Shutdown(shutdownCtx); err != nil {
				slog.Error("failed to shut down server", slog.Any("error", err))
			}
			close(newDumpFilesCh)
			close(newMainFilesCh)
			close(newQuizCodesCh)
			close(errCh)
			wg.Wait()
			return
		}
	}
}

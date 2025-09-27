package config

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

type Config struct {
	HeartbeatIntervalSeconds int
	WatcherPath              string
	SchoolPath               string
	WatcherIntervalSeconds   int
	NumParseFileWorkers      int
	GeminiApiKey             string
}

func New() (*Config, error) {
	cfg := defaultConfig()

	heartbeatIntervalSecondsStr, ok := os.LookupEnv("CURAITOR_HEARTBEAT_INTERVAL_SECONDS")
	if ok {
		heartbeatIntervalSeconds, err := strconv.Atoi(heartbeatIntervalSecondsStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse CURAITOR_HEARTBEAT_INTERVAL_SECONDS")
		}
		cfg.HeartbeatIntervalSeconds = heartbeatIntervalSeconds
	}

	watcherPath, ok := os.LookupEnv("CURAITOR_WATCHER_PATH")
	if ok {
		cfg.WatcherPath = watcherPath
	}

	schoolPath, ok := os.LookupEnv("CURAITOR_SCHOOL_PATH")
	if ok {
		cfg.SchoolPath = schoolPath 
	}

	watcherIntervalSecondsStr, ok := os.LookupEnv("CURAITOR_WATCHER_INTERVAL_SECONDS")
	if ok {
		watcherIntervalSeconds, err := strconv.Atoi(watcherIntervalSecondsStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse CURAITOR_WATCHER_INTERVAL_SECONDS: %w", err)
		}
		cfg.WatcherIntervalSeconds = watcherIntervalSeconds
	}

	numExtractWorkersStr, ok := os.LookupEnv("CURAITOR_NUM_EXTRACT_WORKERS")
	if ok {
		numExtractWorkers, err := strconv.Atoi(numExtractWorkersStr)
		if err != nil {
			return nil, fmt.Errorf("failed to parse CURAITOR_NUM_EXTRACT_WORKERS: %w", err)
		}
		cfg.NumParseFileWorkers = numExtractWorkers
	}

	geminiApiKey, ok := os.LookupEnv("GEMINI_API_KEY")
	if ok {
		cfg.GeminiApiKey = geminiApiKey
	} else {
		return nil, errors.New("GEMINI_API_KEY must be provided")
	}

	slog.Info("config loaded")

	return cfg, nil
}

func defaultConfig() *Config {
	return &Config{
		HeartbeatIntervalSeconds: 60,
		WatcherPath:              "./dump",
		SchoolPath:               "./school", 
		WatcherIntervalSeconds:   5,
		NumParseFileWorkers:      5,
	}
}

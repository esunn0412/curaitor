package data

import (
	"curaitor/internal/model"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"sync"
)

type Edges struct {
	Mu    *sync.Mutex
	Edges []model.Edge
}

func LoadEdges() (*Edges, error) {
	edges, err := loadEdges()

	if err != nil {
		return nil, fmt.Errorf("failed to load edges: %w", err)
	}
	slog.Info("Edges loaded")

	return &Edges{
		Mu:    &sync.Mutex{},
		Edges: edges,
	}, nil

}

func loadEdges() ([]model.Edge, error) {
	var edges []model.Edge

	data, err := os.ReadFile("edges.json")
	if err != nil {
		if os.IsNotExist(err) {
			return edges, nil
		}

		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	if len(data) == 0 {
		return edges, nil
	}

	if err := json.Unmarshal(data, &edges); err != nil {
		return nil, fmt.Errorf("failed to unmarshal edges: %w", err)
	}
	return edges, nil
}

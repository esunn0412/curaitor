package data

import (
	"curaitor/internal/model"
	"sync"
)

type BackLinks struct {
	Mu    *sync.Mutex
	Nodes []model.Node
	Edges []model.Edge
}

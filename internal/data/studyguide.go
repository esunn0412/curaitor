package data

import (
	"curaitor/internal/model"
	"sync"
)

type StudyGuides struct {
	mu     *sync.Mutex
	guides []model.StudyGuide
}
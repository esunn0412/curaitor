package data

import (
	"curaitor/internal/model"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"sync"
)

// this is where we should handle quiz data
type Quiz struct {
	Mu      *sync.Mutex
	Quizzes []model.QuizInfo
}

func LoadQuiz() (*Quiz, error) {
	quizzes, err := loadQuizzes()
	if err != nil {
		return nil, fmt.Errorf("failed to load quizzes: %w", err)
	}

	slog.Info("Quiz loaded")

	return &Quiz{
		Mu:      &sync.Mutex{},
		Quizzes: quizzes,
	}, nil
}

func loadQuizzes() ([]model.QuizInfo, error) {
	var quizzes []model.QuizInfo

	data, err := os.ReadFile("quizzes.json")
	// handle empty json file
	if len(data) == 0 {
		return quizzes, nil
	}
	if err != nil {
		// return without error if courses.json doesn't exist
		if os.IsNotExist(err) {
			return quizzes, nil
		}
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	if err := json.Unmarshal(data, &quizzes); err != nil {
		return nil, fmt.Errorf("failed to unmarshal quizzes: %w", err)
	}

	return quizzes, nil
}

func (q *Quiz) Add(quiz model.QuizInfo) {
	q.Mu.Lock()
	defer q.Mu.Unlock()
	q.Quizzes = append(q.Quizzes, quiz)
}

func (q *Quiz) Save() error {
	q.Mu.Lock()
	defer q.Mu.Unlock()
	data, err := json.MarshalIndent(q.Quizzes, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal quizzes: %w", err)
	}

	if err := os.WriteFile("quizzes.json", data, 0644); err != nil {
		return fmt.Errorf("failed to write quizzes file: %w", err)
	}

	return nil
}

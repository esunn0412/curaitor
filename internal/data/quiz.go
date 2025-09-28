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

// func loadQaPairs() ([]model.QaPairsInfo, error) {
// 	var qaPairs []model.QaPairsInfo

// 	data, err := os.ReadFile("quizzes.json")
// 	if err != nil {
// 		if os.IsNotExist(err) {
// 			return qaPairs, nil
// 		}
// 		return nil, fmt.Errorf("failed to read file: %w", err)
// 	}

// 	if err := json.Unmarshal(data, &qaPairs); err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal qaPairss: %w", err)
// 	}

// 	return qaPairs, nil
// }

// load the quizzes
// func loadQuiz(code string) ([]model.QaPairsInfo, error) {
// 	var quiz []model.QuestionInfo
// 	const quizzesDir = "quizzes" // TODO: add to env

// 	dir_entries, err := os.ReadDir(quizzesDir)
// 	if err != nil {
// 		// return without error
// 		if os.IsNotExist(err) {
// 			return quiz, nil
// 		}
// 		return nil, fmt.Errorf("failed to read directory: %w", err)
// 	}

// 	// Check if quiz exists for a course
// 	for _, dir_entry := range dir_entries {
// 		if dir_entry.Name() == code {
// 			data, err := os.ReadFile("quizzesDir/" + dir_entry.Name())
// 			if err != nil {
// 				return nil, fmt.Errorf("failed to read quiz: %w", err)
// 			}

// 			// Get the quiz file
// 			if err := json.Unmarshal(data, &quiz); err != nil {
// 				return nil, fmt.Errorf("failed to unmarshal quiz: %w", err)
// 			}
// 		}
// 	}

// 	return quiz, nil
// }

func (q *Quiz) Add(quiz model.QuizInfo) {
	q.Mu.Lock()
	defer q.Mu.Unlock()
	q.Quizzes = append(q.Quizzes, quiz)
	slog.Info("Quiz added")
}

func (q *Quiz) Save() error {
	q.Mu.Lock()
	data, err := json.MarshalIndent(q.Quizzes, "", "  ")
	q.Mu.Unlock()

	if err != nil {
		return fmt.Errorf("failed to marshal quizzes: %w", err)
	}

	if err := os.WriteFile("quizzes.json", data, 0644); err != nil {
		return fmt.Errorf("failed to write quizzes file: %w", err)
	}

	return nil
}

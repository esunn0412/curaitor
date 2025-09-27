package gemini

import (
	"context"
	"curaitor/internal/config"
	"curaitor/internal/data"
	"curaitor/internal/model"
	"encoding/json"
	"fmt"
	"io/fs"
	"log/slog"
	"path/filepath"
	"sync"

	"google.golang.org/genai"
)

func GenerateQuizWorker(cfg *config.Config, ctx context.Context, wg *sync.WaitGroup, quizzes *data.Quiz, newQuizCodesCh <-chan string, errCh chan<- error) {
	defer wg.Done()
	var genai_config = &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		ResponseSchema: &genai.Schema{
			Type: genai.TypeArray,
			Items: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"question": {Type: genai.TypeString},
					"choices":  {Type: genai.TypeArray, Items: &genai.Schema{Type: genai.TypeString}},
					"answer":   {Type: genai.TypeInteger},
				},
			},
		},
	}

	for {
		select {
		case code, ok := <-newQuizCodesCh:
			if !ok {
				slog.Info("newQuizCodesCh closed; worker exiting")
				return
			}

			root := filepath.Join(cfg.SchoolPath, code)

			var files []string
			
			err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					slog.Error("Error retrieving file paths.")
				}
				if !d.IsDir() { // Check if it's a file
					files = append(files, path)
				}
				return nil
			})
			if err != nil {
				errCh <- fmt.Errorf("failed to retrieve file paths: %w", err)
			}

			slog.Info("Start generating quiz")
			msg, err := prepMessage(generateQuestionPrompt, files...)
			if err != nil {
				errCh <- fmt.Errorf("failed to prepare message for Gemini: %w", err)
				continue
			}

			res, err := sendMessage(cfg, ctx, msg, genai_config)
			if err != nil {
				errCh <- fmt.Errorf("failed to send message to Gemini: %w", err)
				continue
			}

			qaPairs := []model.Question{}
			if err := json.Unmarshal([]byte(res), &qaPairs); err != nil {
				errCh <- fmt.Errorf("failed to unmarshal Gemini response: %w", err)
				continue
			}

			// iniatilize a quiz struct
			var quiz model.QuizInfo
			quiz.Code = code
			quiz.QaPairs = qaPairs

			// TODO: Handle the deletion of previous quiz
			quizzes.Add(quiz)

			if err := quizzes.Save(); err != nil {
				errCh <- fmt.Errorf("failed to save courses: %w", err)
			}

		case <-ctx.Done():
			slog.Info("worker done")
			return
		}
	}
}

const generateQuestionPrompt = `
You are given a set of input documents about a college course. 
Your task is to carefully read and understand the content, 
then generate 10 multiple choice questions with 4 answer choices with one correct answer.
The question should capture the most important and relevant concepts from the documents. 
The response schema is as follows:
- "question":
	- Try to ask about concepts that can be confusing to students
- "choices":
	- Choices the user should choose from
- "answer:
	- The correct answer value should be the array index of correct choices
	- Try to randomize the correct answer index, so questions don't have a pattern of answers
`

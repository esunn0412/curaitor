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
			Type: genai.TypeObject, 
			Properties: map[string]*genai.Schema{
				"name":  {Type: genai.TypeString},
				"code":  {Type: genai.TypeString},
				"numFiles": {Type: genai.TypeInteger},
				"qaPairs": {
					Type: genai.TypeArray,
					Items: &genai.Schema{
						Type: genai.TypeObject,
						Properties: map[string]*genai.Schema{
							"question": {Type: genai.TypeString},
							"choices": {Type: genai.TypeArray, Items: &genai.Schema{Type: genai.TypeString}},
							"answer":   {Type: genai.TypeString},
						},
					},
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

			slog.Info("trying to walk dir", slog.String("path", root))

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

			slog.Info(res)

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
then generate 10 question-and-answer pairs 
that capture the most important and
relevant concepts from the documents. 
Store each question and corresponding answer strictly in a json object, 
and return only an array of the 10 json object with valid fields:
Do not start and end with '''json .. '''
- "question":
	- Must be a string 
	- Try to ask about concepts that can be confusing to beginner
- "answer:
	- Must be a string
	- must corresponds to the question in the same json object
ex: [{"question":"Calculate the triangle's area with sides 4 and 5cm.","answer":"10"}, {"question":"When is Abraham Lincoln's birthday?","answer":"February 12th, 1809"} ] 
`

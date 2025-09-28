package gemini

import (
	"context"
	"curaitor/internal/config"
	"curaitor/internal/data"
	"curaitor/internal/model"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"sync"

	"google.golang.org/genai"
)

func ParseMainFileWorker(cfg *config.Config, ctx context.Context, wg *sync.WaitGroup, caches *data.CachedFiles, newMainFilesCh <-chan string, errCh chan<- error, fileEdgeCh chan<- string) {
	defer wg.Done()
	config := &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json", 
		ResponseSchema: &genai.Schema{
			Type: genai.TypeArray, 
			Items: &genai.Schema{
				Type: genai.TypeObject, 
				Properties: map[string]*genai.Schema{
					"course_code": {Type: "string"}, 
					"markdown": {Type: "string"}, 
				},
			},
		},
	}

	for {
		select {
		case file := <-newMainFilesCh:
			c, err := os.ReadFile(file)
			if err != nil {
				errCh <- err
				continue
			}

			if len(c) == 0 {
				slog.Warn("empty file, skipping", slog.String("file", file))
				continue
			}

			caches.Add(model.CachedFile{
				FilePath: file,
				Content:  c,
			})
			if err := caches.Save(); err != nil {
				errCh <- fmt.Errorf("failed to save cache: %w", err)
			}

			fileEdgeCh <- file
			msg, err := prepMessageFromCache(parseMainFilePrompt, caches)
			if err != nil {
				errCh <- fmt.Errorf("failed to prepare message for Gemini: %w", err)
				continue
			}

			res, err := sendMessage(cfg, ctx, msg, config)
			if err != nil {
				errCh <- fmt.Errorf("failed to send message to Gemini: %w", err)
				continue
			}

			var studyGuides []model.StudyGuide
			mu := &sync.Mutex{}
			mu.Lock()
			
			if err := json.Unmarshal([]byte(res), &studyGuides); err != nil {
				errCh <- fmt.Errorf("failed to unmarshal Gemini response: %w", err)
				continue
			}
			
			for _, guide := range studyGuides {
				studyGuidePath := filepath.Join(cfg.SchoolPath, guide.CourseCode, "STUDY_GUIDE.md")
				if err := os.WriteFile(studyGuidePath, []byte(guide.Markdown), 0644); err != nil {
					errCh <- fmt.Errorf("failed to write study guide: %w", err)
					continue
				}
				slog.Info("study guide generated", slog.String("path", studyGuidePath))
			}
			mu.Unlock()
		case <-ctx.Done():
			slog.Info("ParseMainFileWorker done")
			return
		}
	}
}

const parseMainFilePrompt = `
	You are given multiple course files. Each file is represented by:
	- A file path
	- Its full content

	Your task is to generate **study guides** in Markdown, grouped by course code.

	### Output format
	Return ONLY valid JSON. Do not include explanations or extra text.  
	The JSON must be an array of objects with this schema:

	[
	{
		"course_code": "string",   // e.g. "CS101"
		"markdown": "string"       // comprehensive study guide in Markdown
	}
	]

	### Study guide requirements
	- Group files by course code (derived from directory structure or filename if present).
	- Combine related files for the same course into a single study guide.
	- Use Markdown features extensively: headings, bullet points, numbered lists, tables, and fenced code blocks where relevant.
	- Ensure content is well-organized into sections and subsections.
	- Be comprehensive: cover all important concepts, examples, and details found in the files.
	- Do not omit useful information, but avoid duplication.
	- Note that study guides should be focused on the course material, not on syllabus or administrative details.

	### Constraints
	- Do not include anything outside of the JSON array.
	- Ensure the JSON is valid and properly escaped.
	- Each markdown field should be a full Markdown document ready to save as a markdown file.
	- Do not include a main title at the top of the markdown, as there will be one already. 

	Now process the following files:

`

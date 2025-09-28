package gemini

import (
	"context"
	"curaitor/internal/config"
	"curaitor/internal/data"
	"fmt"
	"log/slog"
	"os"
	"sync"

	"google.golang.org/genai"
)

func GeminiEdgingWorker(cfg *config.Config, ctx context.Context, wg *sync.WaitGroup, edges data.Edges, fileEdgeCh <-chan string, caches *data.CachedFiles, errCh chan<- error) {
	defer wg.Done()
	var genai_config = &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		ResponseSchema: &genai.Schema{
			Type: genai.TypeArray,
			Items: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"to":   {Type: genai.TypeString},
					"from": {Type: genai.TypeString},
				},
				Required: []string{"to", "from"},
			},
		},
	}

	counter := 0

	for {
		select {
		case _, ok := <-fileEdgeCh:
			counter++

			if !ok {
				slog.Info("fileEdgeCh closed; worker exiting")
				return
			}

			if counter%5 != 0 {
				slog.Info("edge formation debounce", slog.Int("count", counter))
				continue
			}

			cache, err := os.ReadFile("cache.json")
			if err != nil {
				errCh <- fmt.Errorf("failed to read cache: %w", err)
			}

			prompt := edgingPrompt + string(cache)

			slog.Info("creating edges", slog.Int("cache_bytes", len(cache)))

			msg, err := prepMessage(prompt)
			if err != nil {
				errCh <- fmt.Errorf("failed to prepare message for Gemini: %w", err)
				continue
			}

			res, err := sendMessage(cfg, ctx, msg, genai_config)
			if err != nil {
				errCh <- fmt.Errorf("failed to send message to Gemini: %w", err)
				continue
			}

			slog.Info("edges created")

			if err := os.WriteFile("edges.json", []byte(res), 0666); err != nil {
				errCh <- fmt.Errorf("failed to save edges: %w", err)
				continue
			}

			// var edgeS []model.Edge
			// if err := json.Unmarshal([]byte(res), &edgeS); err != nil {
			// 	errCh <- fmt.Errorf("failed to unmarshal Gemini response: %w", err)
			// 	continue
			// }

			// edges.Add(edgeS)
			// if err := edges.Save(); err != nil {
			// 	errCh <- fmt.Errorf("failed to save courses: %w", err)
			// }

		case <-ctx.Done():
			slog.Info("GeminiEdgingWorker done")
			return
		}
	}
}

const edgingPrompt = `
You are given a list of {file_path: string, content: string}. These are file paths and their contents of various college courseworks. Those in the same course code path are under
same college course. File paths are mostly constructed as such: '<course-code>/<file-type>/<file-name>'.
You are an excellent study helper that comprehensively understands all the contents,
and make connections between different files. If contents of any two files are related or relevant, you should make a connection beween them.
Represent the connections as a list of {from: string, to: string}, where each field is the file path. This way, we can construct a graph view that spans across
many files, helping the user understand the concepts better.

Follow these rules:
- Do not create duplicate connections. 'from' and 'to' fields are interchangeable, and there should be no double-linked nodes.
- Try to make connection if any two file paths share the same course name
- Do not make connection with a syllabus file

files:

`

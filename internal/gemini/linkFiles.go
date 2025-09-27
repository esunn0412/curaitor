package gemini

import (
	"context"
	"curaitor/internal/config"
	"curaitor/internal/data"
	"fmt"
	"log/slog"
	"sync"

	"google.golang.org/genai"
)

func geminiEdgingWorker(cfg *config.Config, ctx context.Context, wg *sync.WaitGroup, edges data.Edges, fileEdgeCh <-chan string, caches data.CachedFiles, errCh chan<- error) {
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
			},
		},
	}

	for {
		select {
		case file, ok := <-fileEdgeCh:
			if !ok {
				slog.Info("fileEdgeCh closed; worker exiting")
				return
			}

			msg, err := prepMessage(generateQuestionPrompt, files...)
			if err != nil {
				errCh <- fmt.Errorf("failed to prepare message for Gemini: %w", err)
				continue
			}

		}
	}
}

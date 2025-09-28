package gemini

import (
	"context"
	"curaitor/internal/config"
	"curaitor/internal/data"
	"curaitor/internal/model"
	"encoding/json"
	"fmt"
	"log/slog"
	"sync"

	"google.golang.org/genai"
)

func GeminiEdgingWorker(cfg *config.Config, ctx context.Context, wg *sync.WaitGroup, edges data.Edges, fileEdgeCh <-chan string, caches *data.CachedFiles, errCh chan<- error) {
	defer wg.Done()
	var genai_config = &genai.GenerateContentConfig{
		ResponseMIMEType: "application/json",
		ResponseSchema: &genai.Schema{
			Type: genai.TypeArray,
			// If available in your SDK:
			// MaxItems: ptr(int32(3)),
			Items: &genai.Schema{
				Type: genai.TypeObject,
				Properties: map[string]*genai.Schema{
					"to":   {Type: genai.TypeString},
					"from": {Type: genai.TypeString},
				},
				Required: []string{"to", "from"},
				// If supported: AdditionalProperties: ptr(false),
			},
		},
	}
	// var genai_config = &genai.GenerateContentConfig{
	// 	ResponseMIMEType: "application/json",
	// 	ResponseSchema: &genai.Schema{
	// 		Type: genai.TypeArray,
	// 		Items: &genai.Schema{
	// 			Type: genai.TypeObject,
	// 			Properties: map[string]*genai.Schema{
	// 				"to":   {Type: genai.TypeString},
	// 				"from": {Type: genai.TypeString},
	// 			},
	// 		},
	// 	},
	// }

	for {
		select {
		case file, ok := <-fileEdgeCh:
			if !ok {
				slog.Info("fileEdgeCh closed; worker exiting")
				return
			}

			prompt := "You are given: - A target file path " + file + edgingPrompt

			msg, err := prepMessageFromCache(prompt, caches)
			if err != nil {
				errCh <- fmt.Errorf("failed to prepare message for Gemini: %w", err)
				continue
			}

			res, err := sendMessage(cfg, ctx, msg, genai_config)
			if err != nil {
				errCh <- fmt.Errorf("failed to send message to Gemini: %w", err)
				continue
			}

			var edgeS []model.Edge
			if err := json.Unmarshal([]byte(res), &edgeS); err != nil {
				errCh <- fmt.Errorf("failed to unmarshal Gemini response: %w", err)
				continue
			}

			edges.Add(edgeS)
			if err := edges.Save(); err != nil {
				errCh <- fmt.Errorf("failed to save courses: %w", err)
			}

		case <-ctx.Done():
			slog.Info("GeminiEdgingWorker done")
			return
		}
	}
}

const edgingPrompt = `
 - A list of candidate chunks, each with a file_path and text content snippet.

Task:
- Select up to 3 candidate chunks that are most relevant to TARGET.
- If fewer than 3 are relevant, return fewer.
- If none are relevant, return an empty array.
- Output ONLY JSON that matches the provided schema:
  [{"to":"<TARGET>","from":"<candidate_file_path>"}, ...]
- "to" must always be exactly TARGET.
- "from" must be one of the provided candidate file_path values (no new paths).
- Do not include duplicates.

Relevance guideline examples for college content:
- Same course/unit/topic, direct references, shared assignment, cross-referenced slides, or obvious follow-ups.
- Avoid near-duplicates of the same chunk unless they provide new info.
`

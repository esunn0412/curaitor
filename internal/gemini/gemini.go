package gemini

import (
	"context"
	"curaitor/internal/config"
	"fmt"
	"os"

	"github.com/gabriel-vasile/mimetype"
	"google.golang.org/genai"
)

// sendMessage sends given message to Gemini and returns the
// response as string. Use prepMessage to create messages.
func sendMessage(cfg *config.Config, ctx context.Context, msg []*genai.Content, config *genai.GenerateContentConfig) (string, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  cfg.GeminiApiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return "", fmt.Errorf("failed to create Gemini client: %w", err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		msg,
		config,
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %w", err)
	}

	return result.Text(), nil
}

// prepMessage detects the MIME type of given file and creates
// a new Gemini prompt content with files attached
func prepMessage(prompt string, filepaths ...string) ([]*genai.Content, error) {
	data := make([]*genai.Part, len(filepaths))

	for i, f := range filepaths {
		fileBytes, err := os.ReadFile(f)
		if err != nil {
			return nil, fmt.Errorf("failed to read file: %w", err)
		}

		fileType, err := mimetype.DetectFile(f)
		if err != nil {
			return nil, fmt.Errorf("failed to detect file type: %w", err)
		}

		data[i] = &genai.Part{
			InlineData: &genai.Blob{
				MIMEType: fileType.String(),
				Data:     fileBytes,
			},
		}
	}

	parts := append(
		data,
		genai.NewPartFromText(prompt),
	)

	msg := []*genai.Content{
		genai.NewContentFromParts(parts, genai.RoleUser),
	}

	return msg, nil
}

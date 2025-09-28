package gemini

import (
	"context"
	"curaitor/internal/config"
	"curaitor/internal/data"
	"fmt"
	"os"
	"strings"

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

func prepMessageFromCache(prompt string, caches *data.CachedFiles) ([]*genai.Content, error) {
	caches.Mu.Lock()
	defer caches.Mu.Unlock()
	parts := make([]*genai.Part, 0, len(caches.CachedFiles)*2+1)

	for _, cf := range caches.CachedFiles {
		parts = append(
			parts,
			genai.NewPartFromText(fmt.Sprintf("File: %s", cf.FilePath)),
		)

		fileType, err := mimetype.DetectFile(cf.FilePath)
		if fileType.Is("text/plain") {
			continue
		}

		if err != nil {
			return nil, fmt.Errorf("failed to detect file type: %w", err)
		}

		mime := fileType.String()
		if strings.Contains(mime, ";") {
			mime = strings.SplitN(mime, ";", 2)[0] // strip charset
		}

		parts = append(parts, &genai.Part{
			InlineData: &genai.Blob{
				MIMEType: mime,
				Data:     []byte(cf.Content),
			},
		})
	}
	parts = append(parts, genai.NewPartFromText(prompt))
	msg := []*genai.Content{
		genai.NewContentFromParts(parts, genai.RoleUser),
	}
	return msg, nil
}

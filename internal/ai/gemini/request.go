package gemini

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

const URL = "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent"

type GeminiClient struct {
	apiKey     string
	httpClient *http.Client
}

func NewGeminiClient(ctx context.Context) (*GeminiClient, error) {
	httpClient := &http.Client{
		Timeout: 15 * time.Second,
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, errors.New("[VendorIA] - Missing API KEY")
	}

	return &GeminiClient{
		apiKey:     apiKey,
		httpClient: httpClient,
	}, nil
}

func (gc *GeminiClient) Execute(diff, operation string) (string, error) {
	content, err := getPrompt(operation)
	if err != nil {
		return "", nil
	}
	reqBody := Request{
		Contents: []struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		}{
			{
				Parts: []struct {
					Text string `json:"text"`
				}{
					{Text: fmt.Sprintf("%s\n\n%s", content, diff)},
				},
			},
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("[VendorIA] failed to marshal request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, URL+"?key="+gc.apiKey, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("[VendorIA] failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := gc.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("[VendorIA] request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("[VendorIA] unexpected HTTP status: %d", resp.StatusCode)
	}

	var geminiResp Response
	if err := json.NewDecoder(resp.Body).Decode(&geminiResp); err != nil {
		return "", fmt.Errorf("[VendorIA] failed to decode response: %w", err)
	}

	if len(geminiResp.Candidates) == 0 || len(geminiResp.Candidates[0].Content.Parts) == 0 {
		return "", nil
	}

	return strings.TrimSpace(geminiResp.Candidates[0].Content.Parts[0].Text), nil
}

func getPrompt(operation string) (string, error) {
	if operation == "commit" {
		return os.Getenv("GEMINI_PROMPT"), nil
	}

	if operation == "branch" {
		return os.Getenv("GEMINI_BRANCH_PROMPT"), nil
	}

	return "", errors.New("[VendorIA] failed to read env with prompt")
}

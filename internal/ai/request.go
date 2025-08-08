package ai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const URL = "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent"

type AI struct {
	apiKey     string
	httpClient *http.Client
}

func NewIA(httpClient *http.Client) (*AI, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("[VendorIA] - Error to create env")
	}
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, errors.New("[VendorIA] - Missing API KEY")
	}

	return &AI{
		apiKey:     apiKey,
		httpClient: httpClient,
	}, nil
}

func (ai *AI) Execute(diff string) (string, error) {
	content, err := os.ReadFile("./internal/ai/prompt")
	if err != nil {
		return "", fmt.Errorf("[VendorIA] failed to read .prompt file: %w", err)
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

	req, err := http.NewRequest(http.MethodPost, URL+"?key="+ai.apiKey, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("[VendorIA] failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := ai.httpClient.Do(req)
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

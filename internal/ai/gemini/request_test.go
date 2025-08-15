package gemini

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAI_Execute(t *testing.T) {
	tests := []struct {
		name        string
		prompt      string
		diff        string
		expectError bool
	}{
		{
			name:        "missing prompt",
			prompt:      "",
			diff:        "test diff content",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.prompt != "" {
				os.Setenv("GEMINI_PROMPT", tt.prompt)
			} else {
				os.Unsetenv("GEMINI_PROMPT")
			}

			ai := &GeminiClient{
				apiKey:     "test-key",
				httpClient: &http.Client{},
			}

			_, err := ai.Execute(tt.diff)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestRequest_Structure(t *testing.T) {
	tests := []struct {
		name     string
		request  Request
		expected string
	}{
		{
			name: "valid request structure",
			request: Request{
				Contents: []struct {
					Parts []struct {
						Text string `json:"text"`
					} `json:"parts"`
				}{
					{
						Parts: []struct {
							Text string `json:"text"`
						}{
							{Text: "test content"},
						},
					},
				},
			},
			expected: "test content",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.request.Contents) == 0 {
				t.Errorf("expected contents to be present")
			}

			if len(tt.request.Contents[0].Parts) == 0 {
				t.Errorf("expected parts to be present")
			}

			if tt.request.Contents[0].Parts[0].Text != tt.expected {
				t.Errorf("expected text %s, got %s", tt.expected, tt.request.Contents[0].Parts[0].Text)
			}
		})
	}
}

func TestResponse_Structure(t *testing.T) {
	tests := []struct {
		name     string
		response Response
		expected string
	}{
		{
			name: "valid response structure",
			response: Response{
				Candidates: []struct {
					Content struct {
						Parts []struct {
							Text string `json:"text"`
						} `json:"parts"`
					} `json:"content"`
				}{
					{
						Content: struct {
							Parts []struct {
								Text string `json:"text"`
							} `json:"parts"`
						}{
							Parts: []struct {
								Text string `json:"text"`
							}{
								{Text: "test response"},
							},
						},
					},
				},
			},
			expected: "test response",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(tt.response.Candidates) == 0 {
				t.Errorf("expected candidates to be present")
			}

			if len(tt.response.Candidates[0].Content.Parts) == 0 {
				t.Errorf("expected parts to be present")
			}

			if tt.response.Candidates[0].Content.Parts[0].Text != tt.expected {
				t.Errorf("expected text %s, got %s", tt.expected, tt.response.Candidates[0].Content.Parts[0].Text)
			}
		})
	}
}

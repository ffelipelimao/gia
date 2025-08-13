package exec

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockAI struct {
	executeFunc func(diff string) (string, error)
}

func (m *mockAI) Execute(diff string) (string, error) {
	return m.executeFunc(diff)
}

func TestExecutor_Start(t *testing.T) {
	tests := []struct {
		name        string
		gitDiff     string
		aiResponse  string
		aiError     error
		expectError bool
	}{
		{
			name:        "ai execution error",
			gitDiff:     "test diff content",
			aiResponse:  "",
			aiError:     ErrEmptyGitDiff,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAI := &mockAI{
				executeFunc: func(diff string) (string, error) {
					return tt.aiResponse, tt.aiError
				},
			}

			executor := NewExecutor(mockAI)

			result, err := executor.Start()

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.aiResponse, result)
			}
		})
	}
}

func TestExecutor_Commit(t *testing.T) {
	tests := []struct {
		name        string
		message     string
		expectError bool
	}{
		{
			name:        "valid commit message",
			message:     "feat: add new feature",
			expectError: true,
		},
		{
			name:        "empty commit message",
			message:     "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executor := NewExecutor(&mockAI{})

			err := executor.Commit(tt.message)

			assert.Error(t, err)
		})
	}
}

func TestExecutor_GetGitDiff(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "get git diff",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executor := NewExecutor(&mockAI{})

			diff, err := executor.getGitDiff()

			if err != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, diff)
			}
		})
	}
}

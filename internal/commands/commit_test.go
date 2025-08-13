package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCommitCommand(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "create commit command",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := NewCommitCommand()

			assert.NotNil(t, cmd)
			assert.Equal(t, "commit", cmd.Use)
			assert.Len(t, cmd.Aliases, 1)
			assert.Equal(t, "c", cmd.Aliases[0])
			assert.NotEmpty(t, cmd.Short)
			assert.NotEmpty(t, cmd.Long)
		})
	}
}

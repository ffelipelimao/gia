package ai

import (
	"context"

	"github.com/ffelipelimao/gia/internal/ai/bedrock"
	"github.com/ffelipelimao/gia/internal/ai/gemini"
)

func NewDefaultFactory() *Factory {
	f := NewFactory()

	f.Register("gemini", func(ctx context.Context) (Strategy, error) {
		return gemini.NewGeminiClient(ctx)
	})
	f.Register("bedrock", func(ctx context.Context) (Strategy, error) {
		return bedrock.NewBedrockClient(ctx)
	})
	f.Register("default", func(ctx context.Context) (Strategy, error) {
		return gemini.NewGeminiClient(ctx)
	})
	return f
}

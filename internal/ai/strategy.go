package ai

import (
	"context"
	"fmt"
)

type Strategy interface {
	Execute(diff string) (string, error)
}

type StrategyConstructor func(ctx context.Context) (Strategy, error)

type Factory struct {
	strategies map[string]StrategyConstructor
}

func NewFactory() *Factory {
	return &Factory{strategies: make(map[string]StrategyConstructor)}
}

func (f *Factory) Register(name string, constructor StrategyConstructor) {
	f.strategies[name] = constructor
}

func (f *Factory) Create(ctx context.Context, strategyType string) (Strategy, error) {
	constructor, exists := f.strategies[strategyType]
	if !exists {
		return nil, fmt.Errorf("unknown strategy type: %s", strategyType)
	}
	return constructor(ctx)
}

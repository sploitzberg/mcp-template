package store

import (
	"context"

	"github.com/sploitzberg/mcp-template/internal/core/domain"
)

// Dummy implements ports.Store with fixed data for development and tests
// that need predictable responses. Replace with a real DB adapter (e.g. SQLite)
// in [cmd/app/main.go] when ready.
type Dummy struct{}

// NewDummy returns a store that always returns the same items.
func NewDummy() *Dummy {
	return &Dummy{}
}

// ListItems returns static placeholder items.
func (d *Dummy) ListItems(_ context.Context) ([]domain.Item, error) {
	return []domain.Item{
		{ID: "dummy-1", Title: "Example item"},
		{ID: "dummy-2", Title: "Another example"},
	}, nil
}

package mock

import (
	"context"

	"github.com/sploitzberg/mcp-template/internal/core/domain"
	"github.com/sploitzberg/mcp-template/internal/core/ports"
)

// Store is a test double for [ports.Store].
type Store struct {
	ListItemsFunc func(ctx context.Context) ([]domain.Item, error)
}

// ListItems delegates to ListItemsFunc when set; otherwise returns empty slice.
func (m *Store) ListItems(ctx context.Context) ([]domain.Item, error) {
	if m.ListItemsFunc != nil {
		return m.ListItemsFunc(ctx)
	}
	return nil, nil
}

var _ ports.Store = (*Store)(nil)

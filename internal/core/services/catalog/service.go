package catalog

import (
	"context"
	"fmt"

	"github.com/sploitzberg/mcp-template/internal/core/domain"
	"github.com/sploitzberg/mcp-template/internal/core/ports"
)

// Service implements ports.CatalogService using a Store.
type Service struct {
	store ports.Store
}

// NewService wires the catalog service to its store.
func NewService(store ports.Store) *Service {
	return &Service{store: store}
}

// ListItems returns items from the store.
func (s *Service) ListItems(ctx context.Context) ([]domain.Item, error) {
	items, err := s.store.ListItems(ctx)
	if err != nil {
		return nil, fmt.Errorf("list items: %w", err)
	}
	return items, nil
}

var _ ports.CatalogService = (*Service)(nil)

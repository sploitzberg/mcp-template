package ports

import (
	"context"

	"github.com/sploitzberg/mcp-template/internal/core/domain"
)

// Store is a driven port for persisted or external item data.
type Store interface {
	ListItems(ctx context.Context) ([]domain.Item, error)
}

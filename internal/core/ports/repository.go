package ports

import (
	"context"

	"github.com/sploitzberg/go-hexagonal-template/internal/core/domain"
)

// ResourceRepository is a driven port: the core needs persistence but does
// not know the storage (memory, SQLite, etc.). Adapters implement this.
type ResourceRepository interface {
	Save(ctx context.Context, r *domain.Resource) error
	GetByID(ctx context.Context, id string) (*domain.Resource, error)
}

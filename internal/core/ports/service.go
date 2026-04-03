package ports

import (
	"context"

	"github.com/sploitzberg/go-hexagonal-template/internal/core/domain"
)

// ResourceService is a driver port: defines what the core exposes.
// Handlers (HTTP, CLI) depend on this interface, not concrete implementations.
type ResourceService interface {
	Create(ctx context.Context, content string) (*domain.Resource, error)
	GetByID(ctx context.Context, id string) (*domain.Resource, error)
}

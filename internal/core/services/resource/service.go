package resource

import (
	"context"
	"fmt"
	"time"

	"github.com/sploitzberg/go-hexagonal-template/internal/core/domain"
	"github.com/sploitzberg/go-hexagonal-template/internal/core/ports"
)

// Service implements ports.ResourceService (driver port).
// It orchestrates business logic using driven ports (Hasher, ResourceRepository).
type Service struct {
	hasher ports.Hasher
	repo   ports.ResourceRepository
}

// NewService wires the core to its driven ports via dependency injection.
func NewService(hasher ports.Hasher, repo ports.ResourceRepository) *Service {
	return &Service{
		hasher: hasher,
		repo:   repo,
	}
}

// Create hashes the content, persists the resource, and returns it.
func (s *Service) Create(ctx context.Context, content string) (*domain.Resource, error) {
	if content == "" {
		return nil, &ports.ValidationError{Msg: "content cannot be empty"}
	}
	id, err := s.hasher.Hash(content)
	if err != nil {
		return nil, fmt.Errorf("hash content: %w", err)
	}
	r := &domain.Resource{
		ID:        id,
		Content:   content,
		CreatedAt: time.Now().UTC(),
	}
	if err := s.repo.Save(ctx, r); err != nil {
		return nil, fmt.Errorf("save resource: %w", err)
	}
	return r, nil
}

// GetByID retrieves a resource by ID.
func (s *Service) GetByID(ctx context.Context, id string) (*domain.Resource, error) {
	r, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get resource: %w", err)
	}
	return r, nil
}

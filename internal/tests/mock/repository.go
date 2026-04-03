package mock

import (
	"context"
	"sync"

	"github.com/sploitzberg/go-hexagonal-template/internal/core/domain"
	"github.com/sploitzberg/go-hexagonal-template/internal/core/ports"
)

// MockRepository implements ports.ResourceRepository for tests.
type MockRepository struct {
	mu   sync.RWMutex
	byID map[string]*domain.Resource
}

// NewMockRepository returns an in-memory mock repository.
func NewMockRepository() *MockRepository {
	return &MockRepository{
		byID: make(map[string]*domain.Resource),
	}
}

// Save implements ports.ResourceRepository.
func (m *MockRepository) Save(ctx context.Context, r *domain.Resource) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.byID == nil {
		m.byID = make(map[string]*domain.Resource)
	}
	cpy := *r
	m.byID[r.ID] = &cpy
	return nil
}

// GetByID implements ports.ResourceRepository.
func (m *MockRepository) GetByID(ctx context.Context, id string) (*domain.Resource, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	r, ok := m.byID[id]
	if !ok {
		return nil, nil
	}
	cpy := *r
	return &cpy, nil
}

// Ensure MockRepository implements ports.ResourceRepository at compile time.
var _ ports.ResourceRepository = (*MockRepository)(nil)

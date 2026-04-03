package repository

import (
	"context"
	"sync"

	"github.com/sploitzberg/go-hexagonal-template/internal/core/domain"
)

// Memory implements ports.ResourceRepository using an in-memory map.
// Suitable for tests and demo; swap for SQLite/Postgres in production.
type Memory struct {
	mu   sync.RWMutex
	byID map[string]*domain.Resource
}

// NewMemory returns an in-memory ResourceRepository.
func NewMemory() *Memory {
	return &Memory{
		byID: make(map[string]*domain.Resource),
	}
}

// Save implements ports.ResourceRepository.
func (m *Memory) Save(ctx context.Context, r *domain.Resource) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	// Copy to avoid mutations from callers
	cpy := *r
	m.byID[r.ID] = &cpy
	return nil
}

// GetByID implements ports.ResourceRepository.
func (m *Memory) GetByID(ctx context.Context, id string) (*domain.Resource, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	r, ok := m.byID[id]
	if !ok {
		return nil, nil
	}
	cpy := *r
	return &cpy, nil
}

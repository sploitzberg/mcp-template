package mock

import "github.com/sploitzberg/go-hexagonal-template/internal/core/ports"

// MockHasher implements ports.Hasher for tests.
// Returns deterministic "mock-{content}" for easy assertions.
type MockHasher struct {
	HashFunc func(content string) (string, error)
}

// Hash implements ports.Hasher.
func (m *MockHasher) Hash(content string) (string, error) {
	if m.HashFunc != nil {
		return m.HashFunc(content)
	}
	return "mock-" + content, nil
}

// Ensure MockHasher implements ports.Hasher at compile time.
var _ ports.Hasher = (*MockHasher)(nil)

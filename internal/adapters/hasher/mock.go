package hasher

import "fmt"

// Mock returns a deterministic hash for content (e.g. for tests or demo).
// Production adapters would use crypto/sha256. No external deps.
type Mock struct{}

// NewMock returns a mock hasher.
func NewMock() *Mock {
	return &Mock{}
}

// Hash implements ports.Hasher. Returns "mock-" + content (truncated to 24 chars).
func (m *Mock) Hash(content string) (string, error) {
	if content == "" {
		return "", fmt.Errorf("content cannot be empty")
	}
	s := "mock-" + content
	if len(s) > 24 {
		s = s[:24]
	}
	return s, nil
}

package domain

import "time"

// Resource represents a domain entity.
// Domain models are technology-agnostic; no serialization tags.
type Resource struct {
	ID        string
	Content   string
	CreatedAt time.Time
}

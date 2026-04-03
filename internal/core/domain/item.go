package domain

// Item is a minimal entity for catalog-style data from the store.
// No serialization tags; mapping to wire formats lives in adapters.
type Item struct {
	ID    string
	Title string
}

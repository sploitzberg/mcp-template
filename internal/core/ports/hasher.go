package ports

// Hasher is a driven port: the core needs content hashing but does not
// know the algorithm (SHA256, mock, etc.). Adapters implement this.
type Hasher interface {
	Hash(content string) (string, error)
}

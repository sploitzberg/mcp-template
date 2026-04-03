package ports

// ValidationError indicates invalid input from the caller.
// Handlers use this to return 400 instead of 500.
type ValidationError struct {
	Msg string
}

func (e *ValidationError) Error() string {
	return e.Msg
}

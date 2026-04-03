# Code Review Checklist

## Overview
Checklist for reviewing Go code in this hexagonal architecture codebase.

## Hexagonal Compliance
- [ ] Core does not import adapters or infrastructure packages
- [ ] Domain entities have no `json`, `dynamodbav`, or other serialization tags
- [ ] Handlers depend on port interfaces, not concrete services
- [ ] Wiring happens only in `cmd/app/main.go`
- [ ] New ports/adapters follow `cmd/architecture/HEXAGONAL.md`

## Go Conventions
- [ ] All errors checked and wrapped with `fmt.Errorf("context: %w", err)`
- [ ] Imports grouped: stdlib, then internal (blank line between)
- [ ] Exported functions have comments starting with the function name
- [ ] No external deps added unless necessary (template uses stdlib)
- [ ] `go fmt ./...` and `go vet ./...` pass

## Testing
- [ ] New code has unit tests in `internal/tests/unit/`
- [ ] Mocks in `internal/tests/mock/` implement port interfaces
- [ ] Core services testable in isolation
- [ ] `make test` passes

## Functionality
- [ ] Code does what it's supposed to do
- [ ] Edge cases and errors handled
- [ ] No hardcoded secrets or sensitive data

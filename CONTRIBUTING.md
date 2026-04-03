# Contributing to go-hexagonal-template

Thanks for contributing. This document covers how to set up, develop, and submit changes.

## Prerequisites

- Go 1.25.6 or later
- `golangci-lint` (optional, for local linting): `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest`

## Development Setup

```bash
git clone https://github.com/sploitzberg/go-hexagonal-template.git
cd go-hexagonal-template
go mod download
make test
```

## Before You Start

1. **Read the architecture** — [cmd/architecture/HEXAGONAL.md](cmd/architecture/HEXAGONAL.md) describes the hexagonal pattern. Follow it for new features.
2. **Check conventions** — [AGENTS.md](AGENTS.md) covers build commands, code style, and structure.
3. **Open an issue** (optional) — For larger changes, open an issue first to discuss.

## Workflow

1. **Create a branch** from `main`.
2. **Make changes** following the hexagonal layout (domain → ports → service → adapters → wiring).
3. **Run locally**:
   - `make test`
   - `go build ./...`
   - `go fmt ./...`
   - `go vet ./...`
   - `golangci-lint run` (if installed)
4. **Commit** with clear messages.
5. **Push** and open a pull request.

## Pull Request Checklist

- [ ] `make test` passes
- [ ] `go build ./...` succeeds
- [ ] `go fmt ./...` applied
- [ ] `go vet ./...` passes
- [ ] New code follows [cmd/architecture/HEXAGONAL.md](cmd/architecture/HEXAGONAL.md)
- [ ] Core does not import adapters; domain stays pure
- [ ] Unit tests added for new behavior (use mocks in `internal/tests/mock/`)
- [ ] Wiring in `cmd/app/main.go` updated if adding ports/adapters

## Code Style

- **Indentation**: Tabs for Go (as enforced by `gofmt`)
- **Errors**: Always check; wrap with `fmt.Errorf("context: %w", err)`
- **Imports**: Stdlib first, then internal, with blank lines between groups
- **Exported symbols**: Document with comments starting with the symbol name
- **Testing**: Use mocks from `internal/tests/mock/`; keep core testable in isolation

See [.editorconfig](.editorconfig) for editor settings.

## Adding New Features

Follow the order in [cmd/architecture/HEXAGONAL.md](cmd/architecture/HEXAGONAL.md):

1. Domain entity
2. Driven ports
3. Driver port (if new use case)
4. Service implementation
5. Adapters (driven)
6. Handler (driver)
7. Wiring in `cmd/app/main.go`
8. Unit tests and mocks

## Questions?

Open an issue or discuss in a PR.

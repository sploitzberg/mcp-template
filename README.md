# go-hexagonal-template

A minimal Go boilerplate using [hexagonal architecture](https://alistair.cockburn.us/hexagonal-architecture/) (Ports & Adapters). Zero external dependencies—standard library only. Use this to bootstrap new Go services with a clear, testable structure.

## Quick Start

```bash
# Run the HTTP API (listens on :8080)
make run

# Run tests
make test

# Build for current platform
make build

# Build for all platforms (darwin, linux, windows)
make build-all
```

### Try the API

```bash
# Create a resource
curl -X POST http://localhost:8080/resources -H "Content-Type: application/json" -d '{"content":"hello"}'

# Get by ID (use the ID from the create response)
curl http://localhost:8080/resources/mock-hello
```

## Architecture

The project follows hexagonal (Ports & Adapters) layout:

```
internal/
├── core/           # Business logic (no infra deps)
│   ├── domain/     # Entities
│   ├── ports/      # Interfaces (driver + driven)
│   └── services/   # Use cases
├── adapters/       # Implementations
│   ├── handlers/   # HTTP (driver adapter)
│   ├── hasher/     # Driven adapter
│   └── repository/ # Driven adapter
└── tests/
    ├── mock/       # Test doubles
    └── unit/       # Unit tests

cmd/app/main.go     # Dependency injection / wiring
```

- **Core** never imports adapters. Domain stays pure (no `json` or DB tags).
- **Ports** define interfaces; adapters implement them.
- **Wiring** happens only in `cmd/app/main.go`.

See [cmd/architecture/HEXAGONAL.md](cmd/architecture/HEXAGONAL.md) for the full pattern and how to add new features.

## Development

| Command          | Description                     |
| ---------------- | ------------------------------- |
| `make run`       | Run the app                     |
| `make test`      | Run tests                       |
| `make build`     | Build to `bin/app`              |
| `make build-all` | Cross-compile for all platforms |
| `make clean`     | Remove `bin/`                   |
| `make help`      | Show targets                    |

- **Format**: `go fmt ./...`
- **Lint**: `go vet ./...` or `golangci-lint run`
- **Tests with race detector**: `go test -race ./...`

## Documentation

| Document                                                       | Purpose                                                  |
| -------------------------------------------------------------- | -------------------------------------------------------- |
| [cmd/architecture/HEXAGONAL.md](cmd/architecture/HEXAGONAL.md) | Hexagonal pattern, how to add ports/adapters/services    |
| [AGENTS.md](AGENTS.md)                                         | Build commands, conventions, subagents, skills, commands |

## Requirements

- Go 1.25.6+

## License

See [LICENSE](LICENSE).

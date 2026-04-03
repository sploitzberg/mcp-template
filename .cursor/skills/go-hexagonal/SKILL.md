---
name: go-hexagonal
description: Project context for go-hexagonal-template. Use when working on domain logic, ports, adapters, services, or hexagonal architecture.
---

# Go Hexagonal Template

## Layout (Ports & Adapters)

- `cmd/app/` — main binary, dependency injection
- `cmd/architecture/HEXAGONAL.md` — pattern reference (read first when adding features)
- `internal/core/domain/` — pure entities (no json, dynamodbav tags)
- `internal/core/ports/` — interfaces (driven: Hasher, Repository; driver: ResourceService)
- `internal/core/services/` — use cases implementing driver port
- `internal/adapters/handlers/http/` — driver adapter (HTTP)
- `internal/adapters/hasher/`, `internal/adapters/repository/` — driven adapters
- `internal/tests/mock/` — test doubles
- `internal/tests/unit/` — unit tests

Module: `github.com/sploitzberg/go-hexagonal-template` | Go 1.25.6

## Conventions

- Core never imports adapters
- Domain stays pure; serialization in adapters
- Handlers depend on `ports.ResourceService`, not concrete types
- Use mocks from `internal/tests/mock/` for unit testing
- Wire dependencies only in `cmd/app/main.go`
- Prefer standard library; no external deps in template

## Reference

- [AGENTS.md](../../../AGENTS.md) — build commands, style, subagents
- [cmd/architecture/HEXAGONAL.md](../../../cmd/architecture/HEXAGONAL.md) — add domain, ports, adapters, wiring

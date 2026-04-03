---
name: go-hexagonal
description: Project context for mcp-template (hexagonal Go layout). Use when working on domain logic, ports, adapters, services, or hexagonal architecture.
---

# mcp-template (Hexagonal)

## Layout (Ports & Adapters)

- `cmd/app/` — main binary, dependency injection (MCP HTTP/SSE by default; stdio if `MCP_TRANSPORT=stdio`)
- `docs/architecture/HEXAGONAL.md` — pattern reference (read first when adding features)
- `internal/core/domain/` — pure entities (no json, dynamodbav tags)
- `internal/core/ports/` — interfaces (driven: Store; driver: CatalogService)
- `internal/core/services/catalog/` — use cases implementing driver port
- `internal/adapters/handlers/mcp/` — driver adapter (MCP tools)
- `internal/adapters/store/` — driven adapter (dummy data; replace with DB adapter later)
- `internal/tests/mock/` — test doubles
- `internal/tests/unit/` — unit tests

Module: `github.com/sploitzberg/mcp-template` | Go 1.25.6 | MCP SDK: `github.com/modelcontextprotocol/go-sdk` v1.4.1 (`mcp` package)

## Conventions

- Core never imports adapters or `github.com/modelcontextprotocol/go-sdk/mcp`
- Domain stays pure; serialization and MCP schemas live in adapters
- MCP registration depends on `ports.CatalogService`, not concrete types
- Use mocks from `internal/tests/mock/` for unit testing
- Wire dependencies only in `cmd/app/main.go`
- Baseline dependency: [modelcontextprotocol/go-sdk](https://github.com/modelcontextprotocol/go-sdk) (import path `github.com/modelcontextprotocol/go-sdk/mcp`)

## Reference

- [AGENTS.md](../../../AGENTS.md) — build commands, style, subagents
- [docs/architecture/HEXAGONAL.md](../../../docs/architecture/HEXAGONAL.md) — add domain, ports, adapters, wiring

# Scaffold New Feature

## Overview

Set up a new feature following hexagonal architecture. Read `docs/architecture/HEXAGONAL.md` first.

## Order of Work

1. **Domain** — Create `internal/core/domain/<entity>.go`
   - Pure structs, no infra tags

2. **Driven ports** — Add interfaces in `internal/core/ports/`
   - One file per port (e.g. `store.go`)
   - Use `context.Context` in signatures

3. **Driver port** — Extend `internal/core/ports/service.go` (or add a focused driver port file)
   - Add methods MCP tools will call

4. **Service** — Implement in `internal/core/services/<domain>/service.go`
   - Inject driven ports via constructor
   - Implement driver port

5. **Adapters** — Implement driven ports in `internal/adapters/<name>/`
   - Add mocks in `internal/tests/mock/`
   - Add `var _ ports.X = (*Y)(nil)` for compile-time check

6. **MCP tools** — Extend `internal/adapters/handlers/mcp/` (`mcp.AddTool`)
   - Depend on driver port interface

7. **Wiring** — Update `cmd/app/main.go`
   - Instantiate adapters and pass to service; register tools on the MCP server

8. **Tests** — Add in `internal/tests/unit/`
   - Use mocks from `internal/tests/mock/`

## Verification

- [ ] `make test` passes
- [ ] `go build ./...` succeeds
- [ ] Core does not import adapters or `github.com/modelcontextprotocol/go-sdk/mcp`

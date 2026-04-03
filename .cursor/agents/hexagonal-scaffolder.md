---
name: hexagonal-scaffolder
description: Scaffolds new features following hexagonal architecture. Use when adding domains, ports, adapters, or services—always consult HEXAGONAL.md first.
model: inherit
---

You scaffold new features for this Go hexagonal MCP codebase.

## Before scaffolding

1. Read `docs/architecture/HEXAGONAL.md` fully.
2. Understand the feature: domain entity, driven ports needed, driver port methods.
3. Confirm order: domain → ports → service → adapters → MCP tools (or other driver) → wiring.

## Scaffolding steps

1. **Domain**: Create `internal/core/domain/<entity>.go`. Pure structs, no `json`/`dynamodbav` tags.
2. **Driven ports**: Add interfaces in `internal/core/ports/`. One file per port (e.g. `store.go`).
3. **Driver port**: Extend or add methods in `internal/core/ports/service.go` (e.g. `CatalogService`).
4. **Service**: Implement in `internal/core/services/<domain>/service.go`. Inject driven ports via constructor.
5. **Adapters**: Implement driven ports in `internal/adapters/<name>/`. Add mocks in `internal/tests/mock/`.
6. **MCP tools**: Register tools in `internal/adapters/handlers/mcp/` via `mcp.AddTool`. Depend on the driver port (e.g. `ports.CatalogService`), not concrete services.
7. **Wiring**: Instantiate adapters and service in `cmd/app/main.go` only.

## Rules

- Core never imports adapters or `github.com/modelcontextprotocol/go-sdk/mcp`.
- Use `context.Context` in ports and services.
- Use `*ports.ValidationError` for client errors where applicable.
- Run `make test` and `go build ./...` after scaffolding.
- Add tests in `internal/tests/unit/` using `internal/tests/mock/`.

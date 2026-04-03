---
name: go-add-driven-adapter
description: Add a new driven port and adapter to the hexagonal architecture. Use when the core needs a new capability (e.g. cache, external API, different storage).
---

# Add Driven Port & Adapter

Follow `cmd/architecture/HEXAGONAL.md`. Order: port → adapter → mock → wire.

## Steps

1. **Define port** in `internal/core/ports/<name>.go`:
   ```go
   type MyPort interface {
       DoSomething(ctx context.Context, arg string) (string, error)
   }
   ```

2. **Implement adapter** in `internal/adapters/<name>/` (e.g. `internal/adapters/cache/mock.go`)

3. **Add mock** in `internal/tests/mock/<name>.go` for unit tests

4. **Inject** in `internal/core/services/<domain>/service.go` constructor and fields

5. **Wire** in `cmd/app/main.go`: instantiate adapter, pass to `NewService(...)`

## Rules

- Port uses `context.Context`
- Adapter imports `domain` and `ports`; never the reverse
- Add `var _ ports.MyPort = (*MyAdapter)(nil)` in adapter for compile-time check
- Run `make test` and `go build ./...` after changes

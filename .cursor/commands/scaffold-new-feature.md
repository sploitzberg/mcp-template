# Scaffold New Feature

## Overview
Set up a new feature following hexagonal architecture. Read `cmd/architecture/HEXAGONAL.md` first.

## Order of Work

1. **Domain** — Create `internal/core/domain/<entity>.go`
   - Pure structs, no infra tags

2. **Driven ports** — Add interfaces in `internal/core/ports/`
   - One file per port (e.g. `hasher.go`, `repository.go`)
   - Use `context.Context` in signatures

3. **Driver port** — Extend `internal/core/ports/service.go`
   - Add methods the handlers will call

4. **Service** — Implement in `internal/core/services/<domain>/service.go`
   - Inject driven ports via constructor
   - Implement driver port

5. **Adapters** — Implement driven ports in `internal/adapters/<name>/`
   - Add mocks in `internal/tests/mock/`
   - Add `var _ ports.X = (*Y)(nil)` for compile-time check

6. **Handler** — Add to `internal/adapters/handlers/http/`
   - Depend on driver port interface
   - Add routes in `router.go`

7. **Wiring** — Update `cmd/app/main.go`
   - Instantiate adapters and pass to service

8. **Tests** — Add in `internal/tests/unit/`
   - Use mocks from `internal/tests/mock/`

## Verification
- [ ] `make test` passes
- [ ] `go build ./...` succeeds
- [ ] Core does not import adapters

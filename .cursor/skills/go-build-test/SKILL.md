---
name: go-build-test
description: Build, test, and validate Go code. Use when building, running tests, fixing compilation errors, or verifying the application.
---

# Go Build & Test

## Commands (prefer Makefile)

| Action | Command |
|--------|---------|
| Run app | `make run` or `go run ./cmd/app` |
| Build | `make build` |
| Build all platforms | `make build-all` |
| Test all | `make test` or `go test ./...` |
| Test with race | `go test -race ./...` |
| Test single | `go test -run TestName ./internal/tests/unit` |
| Format | `go fmt ./...` |
| Vet | `go vet ./...` |
| Clean | `make clean` |

## Workflow

1. After code changes: run `make test` and `go build ./...`
2. Before committing: `go fmt ./...` and `go vet ./...`
3. Use `go test -race ./...` when debugging concurrency issues

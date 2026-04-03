---
name: go-build-test
description: Build, test, and validate Go code. Use when building, running tests, fixing compilation errors, or verifying the application.
---

# Go Build & Test

## Commands (prefer Makefile)

| Action | Command |
|--------|---------|
| Run MCP server (HTTP/SSE, default) | `make run` — then Cursor `url` `http://127.0.0.1:8081/sse` |
| Run MCP (stdio) | `MCP_TRANSPORT=stdio make run` |
| Build | `make build` |
| Build all platforms | `make build-all` |
| Test all | `make test` or `go test ./...` |
| Test with race | `go test -race ./...` |
| Test single | `go test -run TestName ./internal/tests/unit` |
| Format | `go fmt ./...` |
| Vet | `go vet ./...` |
| Lint (optional) | `golangci-lint run ./...` — config [`.golangci.yml`](../../.golangci.yml) v2; install: `go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest` |
| Clean | `make clean` |

## Workflow

1. After code changes: run `make test` and `go build ./...`
2. Before committing: `go fmt ./...`, `go vet ./...`, and `golangci-lint run ./...` if installed
3. Use `go test -race ./...` when debugging concurrency issues

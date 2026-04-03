# Agent Development Guidelines for mcp-template

## Build & Test Commands

- **Run (HTTP/SSE, default)**: `make run` or `go run ./cmd/app` — listens on `MCP_HTTP_ADDR` (default `:8081`); **Cursor `mcp.json`**: `"url": "http://127.0.0.1:8081/sse"` while the server is running
- **Run (stdio)**: `MCP_TRANSPORT=stdio make run` — subprocess / stdin–stdout; use `command`/`args`/`cwd` in Cursor, not `url`
- **Build**: `make build` or `go build ./...` — Builds to `bin/app`
- **Build All Platforms**: `make build-all` — Cross-compiles for darwin/linux/windows (amd64, arm64)
- **Test All**: `make test` or `go test ./...` — Runs all tests
- **Test Single**: `go test -run TestName ./internal/tests/unit`
- **Test with Race**: `go test -race ./...` — Run tests with race detection
- **Format**: `go fmt ./...`
- **Lint**: `go vet ./...` or `golangci-lint run` (uses [`.golangci.yml`](.golangci.yml) **v2** schema; install [latest golangci-lint](https://golangci-lint.run/welcome/install/))
- **Clean**: `make clean` — Removes `bin/`

## Code Style & Conventions

- **Package Structure**: `cmd/` for executables, `internal/` for private packages. See `docs/architecture/HEXAGONAL.md` for layout.
- **Imports**: Group as stdlib, then internal packages with blank lines between
- **Naming**: camelCase for variables/functions, PascalCase for exported items
- **Error Handling**: Check all errors; wrap with `fmt.Errorf("context: %w", err)`
- **Comments**: Start with function name for exported functions
- **Testing**: Test files end with `_test.go`; use mocks from `internal/tests/mock/`
- **Architecture**: Follow hexagonal (Ports & Adapters). Core never imports adapters or `github.com/modelcontextprotocol/go-sdk/mcp`.

## Architecture Reference

- **Hexagonal Pattern**: See `docs/architecture/HEXAGONAL.md` for ports, adapters, and wiring in `cmd/app/main.go`
- **Adding Features**: driven port → adapter → extend driver port → service → MCP tool registration → wire in main

## Module: github.com/sploitzberg/mcp-template | Go Version: 1.25.6

## Active Technologies

- Go 1.25.6 + [github.com/modelcontextprotocol/go-sdk](https://github.com/modelcontextprotocol/go-sdk) **v1.4.1** (`mcp` package — MCP server; default HTTP/SSE, optional stdio)
- Hexagonal architecture (Ports & Adapters) with mock-based tests

## AI Workflow & Scaffolding

- When scaffolding new features, consult `docs/architecture/HEXAGONAL.md` first
- Add domain → ports → service → adapters → wiring in that order
- Use `internal/tests/mock/` for test doubles; keep core testable in isolation
- Wire dependencies only in `cmd/app/main.go`; no init-time wiring

## Subagents (`.cursor/agents/`)

| Agent                  | Purpose                                                                 | Invoke                  |
| ---------------------- | ----------------------------------------------------------------------- | ----------------------- |
| `verifier`             | Validates completed work; runs `make test`, checks hexagonal compliance | `/verifier`             |
| `debugger`             | Root cause analysis for errors and test failures                        | `/debugger`             |
| `test-runner`          | Runs tests, fixes failures, adds coverage                               | `/test-runner`          |
| `hexagonal-scaffolder` | Scaffolds new features per HEXAGONAL.md                                 | `/hexagonal-scaffolder` |

## Skills (`.cursor/skills/`)

| Skill                   | Purpose                                  | Invoke                   |
| ----------------------- | ---------------------------------------- | ------------------------ |
| `go-build-test`         | Build, test, validate; Makefile commands | `/go-build-test`         |
| `go-hexagonal`          | Project context, layout, conventions     | `/go-hexagonal`          |
| `go-add-driven-adapter` | Add driven port + adapter                | `/go-add-driven-adapter` |
| `go-add-dep`            | Add Go dependencies beyond the MCP SDK   | `/go-add-dep`            |

## Commands (`.cursor/commands/`)

| Command                 | Purpose                                      | Invoke                   |
| ----------------------- | -------------------------------------------- | ------------------------ |
| `run-tests-and-fix`     | Run `make test`, fix failures systematically | `/run-tests-and-fix`     |
| `code-review-checklist` | Hexagonal + Go review checklist              | `/code-review-checklist` |
| `scaffold-new-feature`  | Step-by-step hexagonal scaffolding           | `/scaffold-new-feature`  |
| `verify-complete`       | Verify work is done; run tests, check wiring | `/verify-complete`       |
| `create-pr`             | PR checklist and description template        | `/create-pr`             |

# mcp-template

A minimal Go [Model Context Protocol](https://modelcontextprotocol.io/) server using [hexagonal architecture](https://alistair.cockburn.us/hexagonal-architecture/) (Ports & Adapters). The default transport is **HTTP/SSE** (for clients like Cursor that connect with a **`url`**). **stdio** is optional.

## Quick Start

```bash
# Run the MCP server (HTTP/SSE on :8081 by default — use Cursor mcp.json "url" below)
make run

# Optional: bind address (default :8081)
MCP_HTTP_ADDR=:9090 make run

# stdio instead (subprocess / stdin–stdout; no "url" in Cursor)
MCP_TRANSPORT=stdio make run

# Run tests
make test

# Build for current platform
make build

# Build for all platforms (darwin, linux, windows)
make build-all
```

### Cursor `mcp.json` (HTTP — primary)

Start the server (`make run`), then add an entry that uses **`url`** only (no `command` / `args` for this server):

```json
"mcp-template": {
  "url": "http://127.0.0.1:8081/sse"
}
```

If you change the listen address with `MCP_HTTP_ADDR`, use the same host and port in `url` (for example `http://127.0.0.1:9090/sse`).

### Optional: stdio (spawned process)

Use this when the client runs the binary itself and talks over stdin/stdout. Set `MCP_TRANSPORT=stdio` for the process.

```json
"mcp-template": {
  "command": "/absolute/path/to/mcp-template/bin/app",
  "args": [],
  "cwd": "/absolute/path/to/mcp-template",
  "env": {
    "MCP_TRANSPORT": "stdio"
  }
}
```

Or with `go run` (keep `cwd` at the module root):

```json
"mcp-template": {
  "command": "go",
  "args": ["run", "./cmd/app"],
  "cwd": "/absolute/path/to/mcp-template",
  "env": {
    "MCP_TRANSPORT": "stdio"
  }
}
```

For the newer **streamable HTTP** transport (`NewStreamableHTTPHandler`), see the [SDK examples](https://github.com/modelcontextprotocol/go-sdk) — this template uses **SSE** (`NewSSEHandler`) at `/sse`, matching patterns like [nimblex402](https://github.com/nimblex402/nimblex402).

## Architecture

```
internal/
├── core/
│   ├── domain/       # Entities (e.g. Item)
│   ├── ports/        # CatalogService (driver), Store (driven)
│   └── services/     # catalog service
├── adapters/
│   ├── handlers/mcp/ # MCP tools → CatalogService, HTTP/SSE bind
│   └── store/        # Store implementation (dummy data; swap for a DB later)
└── tests/
    ├── mock/         # Test doubles
    └── unit/

cmd/app/main.go       # Wiring: store → service → MCP → HTTP/SSE (default) or stdio
```

- **Core** never imports adapters or the MCP SDK.
- **Wiring** happens only in `cmd/app/main.go`.

See [docs/architecture/HEXAGONAL.md](docs/architecture/HEXAGONAL.md) for the full pattern and how to add features.

## Development

| Command          | Description                                      |
| ---------------- | ------------------------------------------------ |
| `make run`       | MCP server on HTTP/SSE (default `:8081`)         |
| `make test`      | Run tests                                        |
| `make build`     | Build to `bin/app`                               |
| `make build-all` | Cross-compile for all platforms                  |
| `make clean`     | Remove `bin/`                                    |
| `make help`      | Show targets                                     |

- **Format**: `go fmt ./...`
- **Lint**: `go vet ./...` and `golangci-lint run ./...` ([v2](https://golangci-lint.run/welcome/install/) — see [`.golangci.yml`](.golangci.yml))
- **Tests with race detector**: `go test -race ./...`

### Pre-push check (local)

```bash
gofmt -w .
go vet ./...
go test -race ./...
go build ./...
golangci-lint run ./...   # optional if installed; matches CI
```

## Documentation

| Document | Purpose |
| -------- | ------- |
| [docs/architecture/HEXAGONAL.md](docs/architecture/HEXAGONAL.md) | Hexagonal pattern, MCP driver adapter, ports |
| [AGENTS.md](AGENTS.md) | Build commands, conventions, subagents, skills |

## Requirements

- **Go** [1.25.6](https://go.dev/dl/) or newer (see [`go.mod`](go.mod))
- **MCP** — [modelcontextprotocol/go-sdk](https://github.com/modelcontextprotocol/go-sdk) **v1.4.1** (`mcp` package); see [MCP documentation](https://modelcontextprotocol.io/docs/learn/server-concepts)

## License

See [LICENSE](LICENSE).

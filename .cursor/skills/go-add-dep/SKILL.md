---
name: go-add-dep
description: Add a new Go dependency. Use when extending the template with external packages (template uses standard library only by default).
---

# Add Go Dependency

This template uses only the standard library. Add deps when needed (DB driver, HTTP client, etc.).

## Process

1. Add the import in your `.go` file
2. Run `go get <module>@<version>` or `go mod tidy`
3. Run `go build ./...` and `make test` to verify

## Examples

```bash
# Add latest
go get github.com/mattn/go-sqlite3

# Add specific version
go get github.com/gin-gonic/gin@v1.9.1

# Add as indirect (tools)
go get -d golang.org/x/tools/cmd/goimports
```

## Notes

- `go mod tidy` cleans unused deps and adds missing ones
- Prefer standard library when sufficient (this template demonstrates that)
- If adding a DB or HTTP framework, add a driven adapter in `internal/adapters/`

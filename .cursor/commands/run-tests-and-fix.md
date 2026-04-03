# Run All Tests and Fix Failures

## Overview
Run the full test suite and systematically fix any failures. Use `make test` and `go build ./...`.

## Steps

1. **Run test suite**
   - Execute `make test` (or `go test ./...`)
   - Capture output and identify failures
   - Run `go test -race ./...` if concurrency is involved

2. **Analyze failures**
   - Identify the failing package and test
   - Check if mocks in `internal/tests/mock/` need updates
   - Determine if the failure is in core, adapter, or wiring

3. **Fix issues systematically**
   - Fix one failure at a time
   - Preserve test intent; do not weaken assertions
   - Re-run `make test` after each fix
   - Ensure `go build ./...` succeeds

## Verification
- [ ] `make test` passes
- [ ] `go build ./...` succeeds

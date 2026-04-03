# Verify Work Complete

## Overview
Verify that claimed work is actually complete and functional. Be skeptical—test everything.

## Steps

1. **Run build and tests**
   - `go build ./...`
   - `make test`
   - If either fails, the work is incomplete

2. **Check implementation exists**
   - Confirm files were created/updated as claimed
   - Verify wiring in `cmd/app/main.go` includes new adapters
   - Ensure handlers are registered in the router

3. **Verify hexagonal compliance**
   - Core does not import `internal/adapters/`
   - Domain has no serialization tags
   - Ports define interfaces; adapters implement them

4. **Report**
   - **Verified**: What passed
   - **Incomplete**: Claims not reflected in code
   - **Broken**: Build or test failures
   - **Violations**: Hexagonal or style issues

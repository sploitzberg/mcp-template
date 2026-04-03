# Create PR

## Overview
Create a well-structured pull request for changes to go-hexagonal-template.

## Pre-PR Checklist
- [ ] `make test` passes
- [ ] `go build ./...` succeeds
- [ ] `go fmt ./...` and `go vet ./...` run
- [ ] New code follows `cmd/architecture/HEXAGONAL.md`
- [ ] Core does not import adapters

## PR Description Template

**Summary**  
Brief description of changes.

**Changes**
- 
- 

**Architecture impact**  
(If adding ports/adapters: which driven port, which adapter, wiring in `cmd/app/main.go`)

**Testing**
- [ ] Unit tests added/updated in `internal/tests/unit/`
- [ ] Mocks added in `internal/tests/mock/` if new port

**Breaking changes**  
(None / describe)

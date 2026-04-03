# Hexagonal Architecture (Ports & Adapters)

This document describes the hexagonal architecture pattern used in this codebase. Follow it when creating new features, adapters, ports, or services. The pattern keeps business logic isolated from infrastructure.

---

## Core Concepts

- **Core**: The center of the hexagon. Contains domain models and business logic. Has no dependency on HTTP, databases, or external systems.
- **Driver actors**: Trigger communication with the core (e.g. HTTP clients, CLI users). They call into the core.
- **Driven actors**: The core calls them (e.g. databases, external APIs). They provide capabilities the core needs.
- **Driver port**: Interface that defines what the core exposes. Handlers depend on this.
- **Driven port**: Interface that defines what the core needs. Adapters implement this.
- **Driver adapter**: Transforms external requests (HTTP, CLI) into core service calls.
- **Driven adapter**: Implements a driven port for a specific technology (e.g. SQLite, Redis, mock).

---

## Directory Structure

```
internal/
├── core/                    # Business logic (no infra deps)
│   ├── domain/               # Entities and value objects
│   │   └── <entity>.go
│   ├── ports/                # Interfaces (driver + driven)
│   │   ├── <driven>.go       # What the core needs (Hasher, Repository)
│   │   ├── <driver>.go       # What the core exposes (Service)
│   │   └── errors.go         # Shared error types (ValidationError, etc.)
│   └── services/
│       └── <domain>/
│           └── service.go    # Implements driver port, uses driven ports
├── adapters/
│   ├── handlers/             # Driver adapters (HTTP, CLI)
│   │   └── http/
│   │       ├── handler.go
│   │       └── router.go
│   ├── <driven_name>/        # Driven adapters (one per port)
│   │   └── mock.go           # or sqlite.go, redis.go, etc.
│   └── repository/
│       └── memory.go
└── tests/
    ├── mock/                 # Test doubles implementing ports
    │   ├── hasher.go
    │   └── repository.go
    └── unit/
        └── <service>_test.go

cmd/
├── app/
│   └── main.go               # Wire adapters to ports (DI)
└── architecture/
    └── HEXAGONAL.md          # This file
```

---

## How to Add New Components

### 1. Domain Entity

Create `internal/core/domain/<entity>.go`. Keep it pure: no `json`, `dynamodbav`, or other infra tags.

```go
package domain

import "time"

type Resource struct {
    ID        string
    Content   string
    CreatedAt time.Time
}
```

### 2. Driven Port (What the Core Needs)

Create `internal/core/ports/<name>.go`. Define the interface the core will depend on.

```go
package ports

type Hasher interface {
    Hash(content string) (string, error)
}
```

For repositories:

```go
package ports

import (
    "context"
    "github.com/.../internal/core/domain"
)

type ResourceRepository interface {
    Save(ctx context.Context, r *domain.Resource) error
    GetByID(ctx context.Context, id string) (*domain.Resource, error)
}
```

### 3. Driver Port (What the Core Exposes)

Create or extend `internal/core/ports/service.go`. Handlers will depend on this interface.

```go
package ports

import (
    "context"
    "github.com/.../internal/core/domain"
)

type ResourceService interface {
    Create(ctx context.Context, content string) (*domain.Resource, error)
    GetByID(ctx context.Context, id string) (*domain.Resource, error)
}
```

### 4. Core Service

Create `internal/core/services/<domain>/service.go`. It implements the driver port and injects driven ports via constructor.

```go
package resource

import (
    "context"
    "fmt"
    "time"
    "github.com/.../internal/core/domain"
    "github.com/.../internal/core/ports"
)

type Service struct {
    hasher ports.Hasher
    repo   ports.ResourceRepository
}

func NewService(hasher ports.Hasher, repo ports.ResourceRepository) *Service {
    return &Service{hasher: hasher, repo: repo}
}

func (s *Service) Create(ctx context.Context, content string) (*domain.Resource, error) {
    if content == "" {
        return nil, &ports.ValidationError{Msg: "content cannot be empty"}
    }
    id, err := s.hasher.Hash(content)
    if err != nil {
        return nil, fmt.Errorf("hash content: %w", err)
    }
    r := &domain.Resource{ID: id, Content: content, CreatedAt: time.Now().UTC()}
    if err := s.repo.Save(ctx, r); err != nil {
        return nil, fmt.Errorf("save resource: %w", err)
    }
    return r, nil
}

func (s *Service) GetByID(ctx context.Context, id string) (*domain.Resource, error) {
    return s.repo.GetByID(ctx, id)
}
```

Rules for the core service:
- Depend only on `ports` interfaces, never on concrete adapters.
- Use `context.Context` for cancellation and tracing.
- Return domain types. Use `*ports.ValidationError` for client errors.

### 5. Driven Adapter

Implement the driven port in `internal/adapters/<name>/`. Example: mock hasher.

```go
package hasher

import "fmt"

type Mock struct{}

func NewMock() *Mock { return &Mock{} }

func (m *Mock) Hash(content string) (string, error) {
    if content == "" {
        return "", fmt.Errorf("content cannot be empty")
    }
    return "hash-" + content, nil
}
```

Example: in-memory repository.

```go
package repository

import (
    "context"
    "sync"
    "github.com/.../internal/core/domain"
)

type Memory struct {
    mu   sync.RWMutex
    byID map[string]*domain.Resource
}

func NewMemory() *Memory {
    return &Memory{byID: make(map[string]*domain.Resource)}
}

func (m *Memory) Save(ctx context.Context, r *domain.Resource) error {
    m.mu.Lock()
    defer m.mu.Unlock()
    cpy := *r
    m.byID[r.ID] = &cpy
    return nil
}

func (m *Memory) GetByID(ctx context.Context, id string) (*domain.Resource, error) {
    m.mu.RLock()
    defer m.mu.RUnlock()
    r, ok := m.byID[id]
    if !ok {
        return nil, nil
    }
    cpy := *r
    return &cpy, nil
}
```

### 6. Driver Adapter (HTTP Handler)

Create `internal/adapters/handlers/http/handler.go`. The handler depends on `ports.ResourceService`, not the concrete service.

```go
package http

import (
    "encoding/json"
    "io"
    "net/http"
    "github.com/.../internal/core/domain"
    "github.com/.../internal/core/ports"
)

type Handler struct {
    svc ports.ResourceService
}

func NewHandler(svc ports.ResourceService) *Handler {
    return &Handler{svc: svc}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
    body, _ := io.ReadAll(r.Body)
    var req struct{ Content string `json:"content"` }
    if err := json.Unmarshal(body, &req); err != nil {
        writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json"})
        return
    }
    res, err := h.svc.Create(r.Context(), req.Content)
    if err != nil {
        var ve *ports.ValidationError
        if errors.As(err, &ve) {
            writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
            return
        }
        writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
        return
    }
    writeJSON(w, http.StatusCreated, res)
}
```

Use separate DTOs for JSON (e.g. `createRequest`, `createResponse`). Do not expose domain structs directly if they need different shapes.

### 7. Wiring in `cmd/app/main.go`

All dependency injection happens here. Choose which adapters to use and connect them to the core.

```go
package main

import (
    "log"
    "net/http"
    httphandler "github.com/.../internal/adapters/handlers/http"
    "github.com/.../internal/adapters/hasher"
    "github.com/.../internal/adapters/repository"
    "github.com/.../internal/core/services/resource"
)

func main() {
    hasherAdapter := hasher.NewMock()
    repoAdapter := repository.NewMemory()

    svc := resource.NewService(hasherAdapter, repoAdapter)
    h := httphandler.NewHandler(svc)
    mux := httphandler.Router(h)

    log.Printf("listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
```

To switch infrastructure (e.g. SQLite instead of memory), change only this file:
replace `repository.NewMemory()` with `repository.NewSQLite(db)`.

### 8. Test Mocks

Create mocks in `internal/tests/mock/` for unit tests. They implement the same port interfaces.

```go
package mock

import "github.com/.../internal/core/ports"

type MockHasher struct {
    HashFunc func(content string) (string, error)
}

func (m *MockHasher) Hash(content string) (string, error) {
    if m.HashFunc != nil {
        return m.HashFunc(content)
    }
    return "mock-" + content, nil
}

var _ ports.Hasher = (*MockHasher)(nil)  // compile-time check
```

Unit test the core service with mocks:

```go
func TestResourceService_Create(t *testing.T) {
    hasher := &mock.MockHasher{}
    repo := mock.NewMockRepository()
    svc := resource.NewService(hasher, repo)
    ctx := context.Background()

    got, err := svc.Create(ctx, "hello")
    if err != nil {
        t.Fatalf("Create: %v", err)
    }
    if got.ID != "mock-hello" {
        t.Errorf("got id %q, want mock-hello", got.ID)
    }
}
```

---

## Flow Summary

```
HTTP Request
    → Handler (driver adapter) parses request
    → Handler calls svc.Create(ctx, content)
    → Service (core) validates, calls hasher.Hash(content)
    → Hasher (driven adapter) returns hash
    → Service calls repo.Save(ctx, resource)
    → Repository (driven adapter) persists
    → Service returns *domain.Resource
    → Handler writes JSON response
```

---

## Rules for LLMs

1. **Core never imports adapters.** Core imports only `domain` and `ports`.
2. **Adapters import core.** Handlers import `ports` and `domain`. Repositories import `domain` and `ports`.
3. **Ports define interfaces.** Driven ports = what the core needs. Driver port = what the core exposes.
4. **Domain is pure.** No `json`, `dynamodbav`, or database tags. Serialization lives in adapters.
5. **Wiring only in `cmd/`.** Dependency injection happens in `main.go`. No `init()` wiring.
6. **Use `context.Context`** in ports and services for cancellation and propagation.
7. **Validation errors** use `*ports.ValidationError` so handlers can return 400.
8. **Test the core in isolation** with `internal/tests/mock` implementations.

---

## Adding a New Driven Port

1. Define interface in `internal/core/ports/<name>.go`
2. Implement in `internal/adapters/<name>/` (e.g. `mock.go`, `sqlite.go`)
3. Add to `NewService()` constructor and fields in `internal/core/services/<domain>/service.go`
4. Create mock in `internal/tests/mock/<name>.go` for unit tests
5. Wire the adapter in `cmd/app/main.go`

## Adding a New Driver Port (Use Case)

1. Add method to driver port in `internal/core/ports/service.go`
2. Implement in `internal/core/services/<domain>/service.go`
3. Add handler method in `internal/adapters/handlers/http/handler.go`
4. Register route in `internal/adapters/handlers/http/router.go`

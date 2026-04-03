package main

import (
	"log"
	"net/http"

	httphandler "github.com/sploitzberg/go-hexagonal-template/internal/adapters/handlers/http"
	"github.com/sploitzberg/go-hexagonal-template/internal/adapters/hasher"
	"github.com/sploitzberg/go-hexagonal-template/internal/adapters/repository"
	"github.com/sploitzberg/go-hexagonal-template/internal/core/services/resource"
)

func main() {
	// Dependency injection: wire adapters to ports
	hasherAdapter := hasher.NewMock()
	repoAdapter := repository.NewMemory()

	svc := resource.NewService(hasherAdapter, repoAdapter)
	h := httphandler.NewHandler(svc)
	mux := httphandler.Router(h)

	addr := ":8080"
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("serve: %v", err)
	}
}

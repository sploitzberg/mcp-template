package unit

import (
	"context"
	"errors"
	"testing"

	"github.com/sploitzberg/go-hexagonal-template/internal/core/ports"
	"github.com/sploitzberg/go-hexagonal-template/internal/core/services/resource"
	"github.com/sploitzberg/go-hexagonal-template/internal/tests/mock"
)

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
	if got.Content != "hello" {
		t.Errorf("got content %q, want hello", got.Content)
	}
}

func TestResourceService_Create_EmptyContent_ReturnsValidationError(t *testing.T) {
	hasher := &mock.MockHasher{}
	repo := mock.NewMockRepository()
	svc := resource.NewService(hasher, repo)
	ctx := context.Background()

	_, err := svc.Create(ctx, "")
	if err == nil {
		t.Fatal("expected error for empty content")
	}
	var ve *ports.ValidationError
	if !errors.As(err, &ve) {
		t.Errorf("expected ValidationError, got %T: %v", err, err)
	}
}

func TestResourceService_GetByID(t *testing.T) {
	hasher := &mock.MockHasher{}
	repo := mock.NewMockRepository()
	svc := resource.NewService(hasher, repo)
	ctx := context.Background()

	created, err := svc.Create(ctx, "test")
	if err != nil {
		t.Fatalf("Create: %v", err)
	}
	got, err := svc.GetByID(ctx, created.ID)
	if err != nil {
		t.Fatalf("GetByID: %v", err)
	}
	if got.ID != created.ID {
		t.Errorf("got id %q, want %q", got.ID, created.ID)
	}
	if got.Content != "test" {
		t.Errorf("got content %q, want test", got.Content)
	}
}

package unit_test

import (
	"context"
	"errors"
	"testing"

	"github.com/sploitzberg/mcp-template/internal/core/domain"
	"github.com/sploitzberg/mcp-template/internal/core/services/catalog"
	"github.com/sploitzberg/mcp-template/internal/tests/mock"
)

func TestCatalogService_ListItems(t *testing.T) {
	ctx := context.Background()
	want := []domain.Item{{ID: "a", Title: "one"}}
	st := &mock.Store{
		ListItemsFunc: func(context.Context) ([]domain.Item, error) {
			return want, nil
		},
	}
	svc := catalog.NewService(st)
	got, err := svc.ListItems(ctx)
	if err != nil {
		t.Fatalf("ListItems: %v", err)
	}
	if len(got) != len(want) || got[0].ID != want[0].ID {
		t.Fatalf("got %+v, want %+v", got, want)
	}
}

func TestCatalogService_ListItems_StoreError(t *testing.T) {
	ctx := context.Background()
	storeErr := errors.New("store unavailable")
	st := &mock.Store{
		ListItemsFunc: func(context.Context) ([]domain.Item, error) {
			return nil, storeErr
		},
	}
	svc := catalog.NewService(st)
	_, err := svc.ListItems(ctx)
	if err == nil {
		t.Fatal("expected error")
	}
	if !errors.Is(err, storeErr) {
		t.Fatalf("expected errors.Is(err, storeErr); got %v", err)
	}
}

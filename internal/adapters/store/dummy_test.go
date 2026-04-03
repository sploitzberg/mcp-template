package store

import (
	"context"
	"testing"
)

func TestDummy_ListItems(t *testing.T) {
	d := NewDummy()
	items, err := d.ListItems(context.Background())
	if err != nil {
		t.Fatalf("ListItems: %v", err)
	}
	if len(items) != 2 {
		t.Fatalf("len(items) = %d, want 2", len(items))
	}
	if items[0].ID != "dummy-1" || items[1].ID != "dummy-2" {
		t.Fatalf("unexpected items: %+v", items)
	}
}

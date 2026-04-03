package mock

import (
	"context"
	"testing"
)

func TestStore_ListItems_DefaultNilFunc(t *testing.T) {
	var m Store
	got, err := m.ListItems(context.Background())
	if err != nil {
		t.Fatalf("ListItems: %v", err)
	}
	if got != nil {
		t.Fatalf("got %v, want nil slice", got)
	}
}

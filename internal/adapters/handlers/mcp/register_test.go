package mcpadapter

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/sploitzberg/mcp-template/internal/core/domain"
	"github.com/sploitzberg/mcp-template/internal/core/services/catalog"
	"github.com/sploitzberg/mcp-template/internal/tests/mock"
)

func TestRegisterTools_list_items(t *testing.T) {
	ctx := context.Background()
	want := []domain.Item{{ID: "t1", Title: "one"}, {ID: "t2", Title: "two"}}
	st := &mock.Store{
		ListItemsFunc: func(context.Context) ([]domain.Item, error) {
			return want, nil
		},
	}
	svc := catalog.NewService(st)

	srv := mcp.NewServer(&mcp.Implementation{Name: "test-server", Version: "0.0.1"}, nil)
	RegisterTools(srv, svc)

	ct, stTrans := mcp.NewInMemoryTransports()
	if _, err := srv.Connect(ctx, stTrans, nil); err != nil {
		t.Fatalf("server.Connect: %v", err)
	}
	client := mcp.NewClient(&mcp.Implementation{Name: "test-client", Version: "0.0.1"}, nil)
	session, err := client.Connect(ctx, ct, nil)
	if err != nil {
		t.Fatalf("client.Connect: %v", err)
	}
	defer session.Close()

	res, err := session.CallTool(ctx, &mcp.CallToolParams{
		Name:      "list_items",
		Arguments: map[string]any{},
	})
	if err != nil {
		t.Fatalf("CallTool: %v", err)
	}
	if res.IsError {
		t.Fatalf("tool returned error")
	}
	var payload struct {
		Items []struct {
			ID    string `json:"id"`
			Title string `json:"title"`
		} `json:"items"`
	}
	raw, err := json.Marshal(res.StructuredContent)
	if err != nil {
		t.Fatalf("marshal structured content: %v", err)
	}
	if err := json.Unmarshal(raw, &payload); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(payload.Items) != len(want) {
		t.Fatalf("got %d items, want %d", len(payload.Items), len(want))
	}
	for i := range want {
		if payload.Items[i].ID != want[i].ID || payload.Items[i].Title != want[i].Title {
			t.Fatalf("item %d: got %+v, want %+v", i, payload.Items[i], want[i])
		}
	}
}

package mcpadapter

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/sploitzberg/mcp-template/internal/core/ports"
)

// RegisterTools adds catalog tools to the MCP server. Core packages must not
// import the MCP SDK; this package is the driver adapter boundary.
func RegisterTools(server *mcp.Server, svc ports.CatalogService) {
	type itemOut struct {
		ID    string `json:"id" jsonschema:"stable item identifier"`
		Title string `json:"title" jsonschema:"human-readable title"`
	}
	type listItemsOut struct {
		Items []itemOut `json:"items" jsonschema:"items returned from the catalog store"`
	}

	mcp.AddTool(server, &mcp.Tool{
		Name:        "list_items",
		Description: "List catalog items from the application's store",
	}, func(ctx context.Context, _ *mcp.CallToolRequest, _ any) (*mcp.CallToolResult, listItemsOut, error) {
		items, err := svc.ListItems(ctx)
		if err != nil {
			return nil, listItemsOut{}, err
		}
		out := listItemsOut{Items: make([]itemOut, 0, len(items))}
		for _, it := range items {
			out.Items = append(out.Items, itemOut{ID: it.ID, Title: it.Title})
		}
		return nil, out, nil
	})
}

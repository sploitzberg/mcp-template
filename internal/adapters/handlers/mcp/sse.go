package mcpadapter

import (
	"context"
	"log"
	"net/http"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// RunSSE serves the MCP server over HTTP using the SDK's SSE transport (same
// pattern as github.com/modelcontextprotocol/go-sdk/mcp.ExampleSSEHandler).
// Clients connect with an HTTP(S) URL, e.g. http://localhost:8081/sse for Cursor.
func RunSSE(ctx context.Context, srv *mcp.Server, addr string) error {
	h := mcp.NewSSEHandler(func(*http.Request) *mcp.Server { return srv }, nil)
	mux := http.NewServeMux()
	mux.Handle("/sse", h)
	mux.Handle("/", h)

	httpSrv := &http.Server{Addr: addr, Handler: mux}
	go func() {
		<-ctx.Done()
		if err := httpSrv.Shutdown(context.Background()); err != nil {
			log.Printf("mcp http shutdown: %v", err)
		}
	}()

	log.Printf("MCP listening (HTTP/SSE) on %s — connect clients to http://%s/sse (or root /)", addr, trimHostPort(addr))
	if err := httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func trimHostPort(addr string) string {
	if len(addr) > 0 && addr[0] == ':' {
		return "localhost" + addr
	}
	return addr
}

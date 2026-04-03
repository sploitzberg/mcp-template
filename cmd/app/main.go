package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	mcpadapter "github.com/sploitzberg/mcp-template/internal/adapters/handlers/mcp"
	"github.com/sploitzberg/mcp-template/internal/adapters/store"
	"github.com/sploitzberg/mcp-template/internal/core/services/catalog"
)

func main() {
	storeAdapter := store.NewDummy()
	svc := catalog.NewService(storeAdapter)

	server := mcp.NewServer(&mcp.Implementation{
		Name:    "mcp-template",
		Version: "0.1.0",
	}, nil)
	mcpadapter.RegisterTools(server, svc)

	transport := strings.ToLower(strings.TrimSpace(os.Getenv("MCP_TRANSPORT")))
	if transport == "stdio" {
		if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
			log.Fatalf("mcp server: %v", err)
		}
		return
	}
	if transport != "" && transport != "http" && transport != "sse" {
		log.Fatalf("mcp server: MCP_TRANSPORT must be stdio, http, sse, or empty (default http/SSE): got %q", transport)
	}

	addr := strings.TrimSpace(os.Getenv("MCP_HTTP_ADDR"))
	if addr == "" {
		addr = ":8081"
	}
	runCtx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	if err := mcpadapter.RunSSE(runCtx, server, addr); err != nil && runCtx.Err() == nil {
		log.Fatalf("mcp server: %v", err)
	}
}

.PHONY: run build build-all test clean help
BINARY  := app
PKG     := github.com/sploitzberg/go-hexagonal-template
MAIN    := ./cmd/app
BINDIR  := bin
LDFLAGS := -ldflags "-s -w"

# Default target
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  run        Run the application"
	@echo "  build      Build for current platform ($(shell go env GOOS)/$(shell go env GOARCH))"
	@echo "  build-all  Build for linux/amd64, linux/arm64, darwin/amd64, darwin/arm64, windows/amd64"
	@echo "  test       Run tests"
	@echo "  clean      Remove build artifacts"
	@echo "  help       Show this help"

run:
	@go run $(MAIN)

build: $(BINDIR)
	@go build $(LDFLAGS) -o $(BINDIR)/$(BINARY) $(PKG)/cmd/app
	@echo "Built $(BINDIR)/$(BINARY)"

build-all: $(BINDIR)
	@set -e; \
	for target in darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64; do \
		os=$${target%%/*}; arch=$${target##*/}; \
		ext=""; [ "$$os" = "windows" ] && ext=".exe"; \
		GOOS=$$os GOARCH=$$arch go build $(LDFLAGS) -o $(BINDIR)/$(BINARY)-$$os-$$arch$$ext $(PKG)/cmd/app; \
	done
	@echo "Built all: $(BINDIR)/$(BINARY)-*"

test:
	@go test ./...

clean:
	@rm -rf $(BINDIR)
	@echo "Cleaned $(BINDIR)"

$(BINDIR):
	@mkdir -p $(BINDIR)

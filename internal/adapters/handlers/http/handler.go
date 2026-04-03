package http

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/sploitzberg/go-hexagonal-template/internal/core/domain"
	"github.com/sploitzberg/go-hexagonal-template/internal/core/ports"
)

// Handler is a driver adapter: transforms HTTP requests into core service calls.
// It depends on ports.ResourceService (driver port), not concrete types.
type Handler struct {
	svc ports.ResourceService
}

// NewHandler returns an HTTP handler wired to the resource service.
func NewHandler(svc ports.ResourceService) *Handler {
	return &Handler{svc: svc}
}

// createRequest is the JSON body for POST /resources.
type createRequest struct {
	Content string `json:"content"`
}

// createResponse is the JSON response for create.
type createResponse struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

// Create handles POST /resources.
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid body"})
		return
	}
	var req createRequest
	if err = json.Unmarshal(body, &req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json"})
		return
	}
	ctx := r.Context()
	res, err := h.svc.Create(ctx, req.Content)
	if err != nil {
		var ve *ports.ValidationError
		if errors.As(err, &ve) {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	writeJSON(w, http.StatusCreated, toCreateResponse(res))
}

// GetByID handles GET /resources/:id.
func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request, id string) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		return
	}
	ctx := r.Context()
	res, err := h.svc.GetByID(ctx, id)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}
	if res == nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
		return
	}
	writeJSON(w, http.StatusOK, toCreateResponse(res))
}

func toCreateResponse(r *domain.Resource) createResponse {
	return createResponse{
		ID:        r.ID,
		Content:   r.Content,
		CreatedAt: r.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("writeJSON: encode error: %v", err)
	}
}

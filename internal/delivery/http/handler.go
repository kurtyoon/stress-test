package http

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"runtime"
	"stress-test/internal/usecase/loadtest"
)

type Handler struct {
	service *loadtest.Service
}

func NewHandler(service *loadtest.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes() {
	http.HandleFunc("/", h.serveDashboard)
	http.HandleFunc("/results", h.getResults)
}

func (h *Handler) serveDashboard(w http.ResponseWriter, r *http.Request) {
	_, b, _, _ := runtime.Caller(0)
	projectRoot := filepath.Join(filepath.Dir(b), "../../../")
	
	dashboardPath := filepath.Join(projectRoot, "web/dashboard.html")
	
	http.ServeFile(w, r, dashboardPath)
}

func (h *Handler) getResults(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.service.GetResults())
} 
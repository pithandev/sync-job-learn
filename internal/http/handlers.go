package http

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/pithandev/sync-job-learn/internal/jobs"
)

type Handler struct {
	store *jobs.Store
}

func NewHandler(store *jobs.Store) *Handler {
	return &Handler{
		store: store,
	}
}

func generateID() string {
	rand.Seed(time.Now().UnixNano())
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 8)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (h *Handler) CreateJob(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	job := &jobs.Job{
		ID:     generateID(),
		Status: jobs.StatusPending,
	}

	h.store.Create(job)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(job)
}

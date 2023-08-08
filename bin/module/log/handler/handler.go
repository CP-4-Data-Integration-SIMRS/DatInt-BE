package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ActivityLog struct {
	Id string
	HC string
	CreatedAt string
}

func InitLogHttpHandler(r *chi.Mux) {
	r.Get("/api/v1", GetLogDataHandler)
}

func GetLogDataHandler(w http.ResponseWriter, r *http.Request) {
	
}




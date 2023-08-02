package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitHttpHandler(r *chi.Mux) {
	r.Get("/api/v1/monitor", GetMonitorDataHandler)
}

func GetMonitorDataHandler(w http.ResponseWriter, r *http.Request) {
	
}




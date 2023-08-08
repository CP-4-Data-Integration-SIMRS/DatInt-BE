package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vier21/simrs-cdc-monitoring/bin/module/monitor/usecase"
)

type httpHandler struct {
	hcUsecase *usecase.HCUsecase
}

func InitMonitorHttpHandler(r *chi.Mux, uc *usecase.HCUsecase) {
	handler := &httpHandler{
		hcUsecase: uc,
	}

	r.Get("/api/v1/monitor", handler.GetMonitorDataHandler)
}

func (h *httpHandler) GetMonitorDataHandler(w http.ResponseWriter, r *http.Request) {
	
}

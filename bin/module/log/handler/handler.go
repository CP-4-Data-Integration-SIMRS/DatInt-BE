package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/vier21/simrs-cdc-monitoring/bin/module/log/usecase"
	"github.com/vier21/simrs-cdc-monitoring/bin/module/log/model"
)

type LogResponse struct {
	Status string           `json:"status"`
	Data   []model.LogData `json:"data"`
}

type httpHandler struct {
	logUsecase usecase.LogUsecase
}

func InitLogHttpHandler(r *chi.Mux, uc usecase.LogUsecase) {
	handler := &httpHandler{
		logUsecase: uc,
	}

	r.Get("/api/v1/logs", handler.GetLogsHandler)
}

func (h *httpHandler) GetLogsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	logs, err := h.logUsecase.GetLogs()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(LogResponse{
			Status: fmt.Sprintf("error fetching data: %s (%s)", err.Error(), http.StatusText(http.StatusInternalServerError)),
			Data:   nil,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(LogResponse{
		Status: fmt.Sprintf("Success (%s)", http.StatusText(http.StatusOK)),
		Data:   logs,
	})
}

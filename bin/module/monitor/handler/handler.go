package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/vier21/simrs-cdc-monitoring/bin/module/monitor/usecase"
)

type MonitorResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type httpHandler struct {
	hcUsecase *usecase.HCUsecase
}

func InitMonitorHttpHandler(r *chi.Mux, uc *usecase.HCUsecase) {
	handler := &httpHandler{
		hcUsecase: uc,
	}

	r.Get("/api/v1/monitor", handler.GetMonitorDataHandler)
	r.Get("/api/v1/{dbname}/monitor", handler.GetDBInfo)
}

func (h *httpHandler) GetMonitorDataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	data, err := h.hcUsecase.GetAllDatabaseInfo()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(MonitorResponse{
			Status: fmt.Sprintf("error fetching data: %s (%s)", err.Error(), strconv.Itoa(http.StatusInternalServerError)),
			Data:   nil,
		})
		return
	}

	if err := json.NewEncoder(w).Encode(MonitorResponse{
		Status: fmt.Sprintf("Success (%s)", strconv.Itoa(http.StatusOK)),
		Data:   data,
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(MonitorResponse{
			Status: fmt.Sprintf("error fetching data: %s (%s)", err.Error(), strconv.Itoa(http.StatusInternalServerError)),
			Data:   nil,
		})
		return
	}
}

func (h *httpHandler) GetDBInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	dbname := chi.URLParam(r, "dbname")

	data, err := h.hcUsecase.GetDBInfo(dbname)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(MonitorResponse{
			Status: fmt.Sprintf("db not found: %s (%s)", err.Error(), strconv.Itoa(http.StatusBadRequest)),
			Data:   nil,
		})
		return
	}

	if err := json.NewEncoder(w).Encode(MonitorResponse{
		Status: fmt.Sprintf("Success (%s)", strconv.Itoa(http.StatusOK)),
		Data:   data,
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(MonitorResponse{
			Status: fmt.Sprintf("error fetching data: %s (%s)", err.Error(), strconv.Itoa(http.StatusInternalServerError)),
			Data:   nil,
		})
		return
	}

}

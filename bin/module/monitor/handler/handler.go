package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/IBM/sarama"
	"github.com/go-chi/chi/v5"
	"github.com/vier21/simrs-cdc-monitoring/bin/module/monitor/model"
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

	r.Get("/api/v1/monitor", handler.GetAllDBNameHandler)
	r.Get("/api/v1/monitor/search", handler.SearchDBNameFromElastic)
	r.Get("/api/v1/{dbname}/monitor", handler.GetDBInfo)
	r.Post("/api/v1/monitor", handler.PushToBroker)

}

func (h *httpHandler) GetMonitorDataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

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
	w.Header().Set("Access-Control-Allow-Origin", "*")

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

func (h *httpHandler) PushToBroker(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	brokersUrl := []string{"localhost:9092"}
	producer, err := ConnectProducer(brokersUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer producer.Close()
	var data []model.DatabaseInfo

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Println(err)
		return
	}

	bytes, _ := json.Marshal(data)

	msg := &sarama.ProducerMessage{
		Topic: "mntr5",
		Value: sarama.StringEncoder(bytes),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", msg.Topic, partition, offset)
}

func (h *httpHandler) GetAllDBNameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	dbname := r.URL.Query().Get("dbname")

	if dbname == "" {
		dbnames, err := h.hcUsecase.GetDbNameFromElastic()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(MonitorResponse{
				Status: fmt.Sprintf("error fetching data: %s (%s)", err.Error(), strconv.Itoa(http.StatusInternalServerError)),
				Data:   nil,
			})
			return
		}

		resp := MonitorResponse{
			Status: fmt.Sprintf("Success (%s)", strconv.Itoa(http.StatusOK)),
			Data:   dbnames,
		}

		if err := json.NewEncoder(w).Encode(resp); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(MonitorResponse{
				Status: fmt.Sprintf("error fetching data: %s (%s)", err.Error(), strconv.Itoa(http.StatusInternalServerError)),
				Data:   nil,
			})
			return
		}
		return

	}
	dbs, err := h.hcUsecase.SearchDBName(dbname)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(MonitorResponse{
			Status: fmt.Sprintf("error fetching data: %s (%s)", err.Error(), strconv.Itoa(http.StatusInternalServerError)),
			Data:   nil,
		})
		return
	}

	resp := MonitorResponse{
		Status: fmt.Sprintf("Success (%s)", strconv.Itoa(http.StatusOK)),
		Data:   dbs,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(MonitorResponse{
			Status: fmt.Sprintf("error fetching data: %s (%s)", err.Error(), strconv.Itoa(http.StatusInternalServerError)),
			Data:   nil,
		})
		return
	}

}

func (h *httpHandler) SearchDBNameFromElastic(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	dbname := r.URL.Query().Get("dbname")

	dbs, err := h.hcUsecase.SearchDBName(dbname)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(MonitorResponse{
			Status: fmt.Sprintf("error fetching data: %s (%s)", err.Error(), strconv.Itoa(http.StatusInternalServerError)),
			Data:   nil,
		})
		return
	}

	resp := MonitorResponse{
		Status: fmt.Sprintf("Success (%s)", strconv.Itoa(http.StatusOK)),
		Data:   dbs,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(MonitorResponse{
			Status: fmt.Sprintf("error fetching data: %s (%s)", err.Error(), strconv.Itoa(http.StatusInternalServerError)),
			Data:   nil,
		})
		return
	}

}

func ConnectProducer(brokersUrl []string) (sarama.SyncProducer, error) {

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

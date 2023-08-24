package kafka

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/vier21/simrs-cdc-monitoring/bin/module/monitor/model"
)

type Producer struct{}

func NewProducer() *Producer {
	return &Producer{}
}

func (p *Producer) SendDataToProducer(data []model.DatabaseInfo) error {
	producerUrl := "http://localhost:3030/api/v1/monitor"

	requestBodyBytes, err := json.Marshal(data)

	if err != nil {
		log.Println("error here")
		return err
	}

	requestBody := bytes.NewReader(requestBodyBytes)
	req, err := http.NewRequest("POST", producerUrl, requestBody)

	if err != nil {
		log.Println("error here")

		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return fmt.Errorf("request failed: %s", err.Error())
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("unable to send data: %s", err.Error())
	}
	return nil
}

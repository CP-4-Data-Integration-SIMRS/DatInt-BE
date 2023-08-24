package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/google/uuid"
)

type LogData struct {
	Healthcare string
	DBName     string
	TBName     string
	Status     string
	DateTime   time.Time
	CreatedAt  time.Time
	RecordId   uuid.UUID
}

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200", // Ganti dengan URL dan port Elasticsearch Anda
		},
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	data := []LogData{
		{
			Healthcare: "Hospital Siloam",
			DBName:     "Database 2",
			TBName:     "PatientInfo",
			Status:     "Success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "Hospital Mayapada ",
			DBName:     "Database 3",
			TBName:     "Appointments",
			Status:     "Success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "Hospital Roem",
			DBName:     "Database 9",
			TBName:     "Invoices",
			Status:     "Failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "Hospital Eldelweis",
			DBName:     "Database 11",
			TBName:     "Invoices",
			Status:     "Failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "Hospital Hasan",
			DBName:     "Database 7",
			TBName:     "PatientName",
			Status:     "Success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		}, {
			Healthcare: "Hospital Hasani",
			DBName:     "Database 10",
			TBName:     "RoomInfo",
			Status:     "Success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "Hospital Internasional",
			DBName:     "Database 1",
			TBName:     "RawatJalan",
			Status:     "Failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
	}

	for _, d := range data {
		body := fmt.Sprintf(`{
			"Healthcare": "%s",
			"DBName": "%s",
			"TBName": "%s",
			"Status": "%s",
			"DateTime": "%s",
			"CreatedAt": "%s",
			"RecordId": "%s"
		}`, d.Healthcare, d.DBName, d.TBName, d.Status, d.DateTime.Format(time.RFC3339), d.CreatedAt.Format(time.RFC3339), d.RecordId)

		req := esapi.IndexRequest{
			Index:      "logindex", // Ganti dengan nama indeks yang sesuai
			DocumentID: d.RecordId.String(),
			Body:       strings.NewReader(body),
			Refresh:    "true",
		}

		res, err := req.Do(context.Background(), client)
		if err != nil {
			log.Printf("Error indexing document: %s", err)
		}
		defer res.Body.Close()

		if res.IsError() {
			log.Printf("Error indexing document: %s", res.String())
		} else {
			log.Printf("Indexed document with ID: %s", d.RecordId.String())
		}
	}
}

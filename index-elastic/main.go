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
	Healthcare string    `json:"healthcare"`
	DBName     string    `json:"dbname"`
	TBName     string    `json:"tbname"`
	Status     string    `json:"status"`
	DateTime   time.Time `json:"dateTime"`
	CreatedAt  time.Time `json:"createdAt"`
	RecordId   uuid.UUID `json:"recordId"`
}

func main() {
	cfg := elasticsearch.Config{
		CloudID: "es-dbt:YXNpYS1zb3V0aGVhc3QxLmdjcC5lbGFzdGljLWNsb3VkLmNvbSRkMjYyNWJjNzY4NjA0ZDM1YTkzOWQyNWU2ZjI0NmJjMCQyMWI3Mjg3MjY2OWY0OTBmOTU3MTk1MjQ4ZGQ3YWNmNg==",
		APIKey:  "SDZ0Nko0b0JyTnVOd2FaWVN1NHI6WHhhZkNYQ1RSb1dtcU0zWUN4YUQxdw==",
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	data := []LogData{
		{
			Healthcare: "rs_siloam",
			DBName:     "database_2",
			TBName:     "patient_info",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_mayapada ",
			DBName:     "database_3",
			TBName:     "appointments",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_roem",
			DBName:     "database_9",
			TBName:     "invoices",
			Status:     "failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_eldelweis",
			DBName:     "database_11",
			TBName:     "invoices",
			Status:     "failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_hasan",
			DBName:     "database_7",
			TBName:     "patient_name",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		}, {
			Healthcare: "rs_hasani",
			DBName:     "database_10",
			TBName:     "room_info",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_internasional",
			DBName:     "database_69",
			TBName:     "rawat_jalan",
			Status:     "failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_nasi_padang",
			DBName:     "database_14",
			TBName:     "rawat_inap",
			Status:     "failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_raos",
			DBName:     "database 17",
			TBName:     "rawat_rujukan",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_nasional",
			DBName:     "database_19",
			TBName:     "rawatan_inap",
			Status:     "failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_bermuda",
			DBName:     "database_21",
			TBName:     "rawat_saja",
			Status:     "failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_bertua",
			DBName:     "database_1",
			TBName:     "rawat_sembuh",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_nonstop",
			DBName:     "database_23",
			TBName:     "rawat_sekarang",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_uks",
			DBName:     "database_25",
			TBName:     "rawat_untuk",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_telkom",
			DBName:     "database_27",
			TBName:     "rawat_opname",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_univ",
			DBName:     "database_30",
			TBName:     "rawat_kos",
			Status:     "failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_militer",
			DBName:     "database_33",
			TBName:     "rawat_intensif",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_umum",
			DBName:     "database_36",
			TBName:     "rawat_vip",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_daerah",
			DBName:     "database_39",
			TBName:     "rawat_klinik",
			Status:     "failed",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
		{
			Healthcare: "rs_gigi",
			DBName:     "database_40",
			TBName:     "rawat_gigi",
			Status:     "success",
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		},
	}

	for _, d := range data {
		body := fmt.Sprintf(`{
			"healthcare": "%s",
			"dbame": "%s",
			"tbname": "%s",
			"status": "%s",
			"dateTime": "%s",
			"createdAt": "%s",
			"recordId": "%s"
		}`, d.Healthcare, d.DBName, d.TBName, d.Status, d.DateTime.Format(time.RFC3339), d.CreatedAt.Format(time.RFC3339), d.RecordId)

		req := esapi.IndexRequest{
			Index:      "search-log", // Ganti dengan nama indeks yang sesuai
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

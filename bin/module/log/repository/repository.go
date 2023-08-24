package repository

import (
	"context"
	"encoding/json"

	//"strings"
	"bytes"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/google/uuid"
	"github.com/vier21/simrs-cdc-monitoring/bin/module/log/model"
	"github.com/vier21/simrs-cdc-monitoring/bin/pkg/elastic"
)

type LogRepositoryInterface interface {
	GetLogs(status, search string) ([]model.LogData, error)
}

type LogRepository struct {
	es *elasticsearch.Client
}

func NewLogRepository() *LogRepository {
	return &LogRepository{
		es: elastic.ESCli,
	}
}

func (lr *LogRepository) GetLogs(status, search string) ([]model.LogData, error) {
	var logs []model.LogData
	var searchBody string

	if status != "" && search != "" {
		// Jika status dan search  tidak kosong
		searchBody = `
		{
			"query": {
				"bool": {
					"must": [
						{
							"match": {
								"Status": "` + status + `"
							}
						},
						{
							"multi_match": {
								"query": "` + search + `",
								"fields": ["Healthcare", "DBName", "TBName"]
							}
						}
					]
				}
			}
		}
		`
	} else if status != "" {
		// Jika status tidak kosong 
		searchBody = `
		{
			"query": {
				"match": {
					"Status": "` + status + `"
				}
			}
		}
		`
	} else if search != "" {
		// Jika search tidak kosong 
		searchBody = `
		{
			"query": {
				"multi_match": {
					"query": "` + search + `",
					"fields": ["Healthcare", "DBName", "TBName"]
				}
			}
		}

		`
	} else {
		// Keduanya kosong
		searchBody = `
		{
			"query": {
				"match_all": {}
			}
		}
		`
	}

	req := esapi.SearchRequest{
		Index: []string{"logindex"},
		Body:  bytes.NewReader([]byte(searchBody)),
	}

	res, err := req.Do(context.Background(), lr.es)
	if err != nil {
		log.Fatalf("Error performing search request: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error response: %s", res.String())
	}

	var response map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	hits := response["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})

		log := model.LogData{
			Healthcare: source["Healthcare"].(string),
			DBName:     source["DBName"].(string),
			TBName:     source["TBName"].(string),
			Status:     source["Status"].(string),
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordID:   uuid.New(),
		}
		logs = append(logs, log)
	}

	return logs, nil
}


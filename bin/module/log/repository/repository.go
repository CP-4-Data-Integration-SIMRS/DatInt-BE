package repository

import (
	"context"
	"encoding/json"
	//"strings"
	"time"
	"log"
	"bytes"


	"github.com/google/uuid"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/vier21/simrs-cdc-monitoring/bin/module/log/model"
	"github.com/vier21/simrs-cdc-monitoring/bin/pkg/elastic"
)

type LogRepositoryInterface interface {
	GetLogs() ([]model.LogData, error)
	
}

type LogRepository struct {
	es *elasticsearch.Client
}

func NewLogRepository() *LogRepository {
	return &LogRepository{
		es: elastic.ESCli,
	}
}

func (lr *LogRepository) GetLogs() ([]model.LogData, error) {
	var logs []model.LogData
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"}, 
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	searchBody := `
	{
		"query": {
			"match_all": {}
		}
	}
	`
	req := esapi.SearchRequest{
		Index: []string{"logindex"}, // Ganti dengan nama indeks Anda
		Body:  bytes.NewReader([]byte(searchBody)),
	}

	res, err := req.Do(context.Background(), es)
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


	// // Your Elasticsearch query
	// query := map[string]interface{}{
	// 	"query": map[string]interface{}{
	// 		"match_all": map[string]interface{}{},
	// 	},
	// }

	
	// jsonStr, err := json.Marshal(query)
	// if err != nil {
	// 	return nil, err
	// }

	
	// res, err := lr.es.Search(
	// 	lr.es.Search.WithContext(context.Background()),
	// 	lr.es.Search.WithIndex("indexlog"), 
	// 	lr.es.Search.WithBody(strings.NewReader(string(jsonStr))),
	// )
	// if err != nil {
	// 	return nil, err
	// }
	// defer res.Body.Close()

	
	// var result map[string]interface{}
	// if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
	// 	return nil, err
	// }
	// fmt.Println(result)
	// hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	// for _, hit := range hits {
	// 	source := hit.(map[string]interface{})["_source"].(map[string]interface{})
	
	// 	// // Handle case where "Status" field is not present
	// 	// statusVal, statusExists := source["Status"].(string)
	// 	// if !statusExists {
	// 	// 	// Provide a default value or take other actions
	// 	// 	statusVal = "success"
	// 	// }
	
	// 	log := model.LogData{
	// 		Healthcare: source["Healthcare"].(string),
	// 		DBName:     source["DBName"].(string),
	// 		TBName:     source["TBName"].(string),
	// 		Status:     source["Status"].(string),
	// 		DateTime:   time.Now(),
	// 		CreatedAt:  time.Now(),
	// 		RecordID:   uuid.New(),
	// 	}
	// 	logs = append(logs, log)
	// }
	

	return logs, nil
}
package repository

import (
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/elastic/go-elasticsearch/v8"
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

	// Your Elasticsearch query
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match_all": map[string]interface{}{},
		},
	}

	
	jsonStr, err := json.Marshal(query)
	if err != nil {
		return nil, err
	}

	
	res, err := lr.es.Search(
		lr.es.Search.WithContext(context.Background()),
		lr.es.Search.WithIndex("test_index"), 
		lr.es.Search.WithBody(strings.NewReader(string(jsonStr))),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
	
		// Handle case where "Status" field is not present
		statusVal, statusExists := source["Status"].(string)
		if !statusExists {
			// Provide a default value or take other actions
			statusVal = ""
		}
	
		log := model.LogData{
			Healthcare: source["Healthcare"].(string),
			DBName:     source["DBName"].(string),
			TBName:     source["TBName"].(string),
			Status:     statusVal,
			DateTime:   time.Now(),
			CreatedAt:  time.Now(),
			RecordId:   uuid.New(),
		}
		logs = append(logs, log)
	}
	

	return logs, nil
}

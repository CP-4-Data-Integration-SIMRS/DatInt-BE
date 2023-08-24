package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/vier21/simrs-cdc-monitoring/bin/module/monitor/model"
	"github.com/vier21/simrs-cdc-monitoring/bin/pkg/elastic"
	"github.com/vier21/simrs-cdc-monitoring/bin/pkg/mysql"
)

type HCRepositoryInterface interface {
	GetAllDB() ([]string, error)
	GetAllTableByDB(dbname string) ([]string, error)
	CountTotalTablesByDBName(dbname string) (int, error)
	CountTableRecords(dbname string, tablename string) (int, error)
	CountNewData(dbname string) (int, error)
	CountDeltaData(dbname string) (int, error)
	FetchNamesFromElasticsearch() ([]string, error)
	SearchDocumentsByDBName(indexName, dbname string) (model.DatabaseInfo, error)
}

type HCRepository struct {
	db      *sqlx.DB
	elastic *elasticsearch.Client
}

func NewHealthCareRepository() *HCRepository {
	return &HCRepository{
		db:      mysql.DB,
		elastic: elastic.ESCli,
	}
}

func (h *HCRepository) GetAllDB() ([]string, error) {
	var dbs []string

	if err := h.db.Select(&dbs, "SHOW DATABASES"); err != nil {
		log.Println(err)
		return []string{}, fmt.Errorf("error get all db %s", err.Error())
	}

	return dbs, nil
}

func (h *HCRepository) CountTotalTablesByDBName(dbname string) (int, error) {
	var count int

	if err := h.db.Get(&count, "SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = ?", dbname); err != nil {
		log.Println(err)
		return 0, fmt.Errorf("error count table %s", err.Error())
	}

	return count, nil
}

func (h *HCRepository) CountTableRecords(dbname string, tablename string) (int, error) {
	totalrecord := 0
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s.%s", dbname, tablename)

	if err := h.db.Get(&totalrecord, query); err != nil {
		log.Println(err)

		return 0, err
	}

	return totalrecord, nil
}

func (h *HCRepository) GetAllTableByDB(dbname string) ([]string, error) {
	var tbnames []string

	if err := h.db.Select(&tbnames, "SELECT TABLE_NAME FROM information_schema.TABLES WHERE TABLE_SCHEMA = ?", dbname); err != nil {
		log.Println(err)
		return []string{}, err
	}

	return tbnames, nil
}

func (h *HCRepository) FetchNamesFromElasticsearch() ([]string, error) {
	indexName := "search-monitor"

	req := esapi.SearchRequest{
		Index: []string{indexName},
		Body:  strings.NewReader(`{"_source": ["dbname"], "query": {"match_all": {}}}`),
	}

	res, err := req.Do(context.Background(), h.elastic)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("error response: %s", res)
	}

	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	hits, ok := result["hits"].(map[string]interface{})["hits"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("hits not found in response")
	}

	var dbNames []string
	for _, hit := range hits {
		source := hit.(map[string]interface{})["_source"].(map[string]interface{})
		dbName := source["dbname"].(string)
		dbNames = append(dbNames, dbName)
	}

	return dbNames, nil
}

func (h *HCRepository) SearchDocumentsByDBName(indexName, dbname string) (model.DatabaseInfo, error) {
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"dbname": dbname,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return model.DatabaseInfo{}, err
	}

	res, err := h.elastic.Search(
		h.elastic.Search.WithContext(context.Background()),
		h.elastic.Search.WithIndex(indexName),
		h.elastic.Search.WithBody(&buf),
		h.elastic.Search.WithTrackTotalHits(true),
		h.elastic.Search.WithPretty(),
	)
	if err != nil {
		return model.DatabaseInfo{}, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return model.DatabaseInfo{}, fmt.Errorf("error response: %s", res.Status())
	}

	var response map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return model.DatabaseInfo{}, err
	}

	totalHits := int(response["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))
	if totalHits > 0 {
		hits := response["hits"].(map[string]interface{})["hits"].([]interface{})
		firstHit := hits[0].(map[string]interface{})
		documentSource := firstHit["_source"].(map[string]interface{})
		jsonData, err := json.Marshal(documentSource)
		if err != nil {
			return model.DatabaseInfo{}, err
		}

		var dbinfo model.DatabaseInfo
		err = json.Unmarshal([]byte(jsonData), &dbinfo)
		if err != nil {
			return model.DatabaseInfo{}, err
		}
		return dbinfo, nil
	}
	return model.DatabaseInfo{}, nil
}
func (h *HCRepository) CountNewData(dbname string) (int, error) {

	return 0, nil
}

func (h *HCRepository) CountDeltaData(dbname string) (int, error) {
	return 0, nil
}

package repository

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/vier21/simrs-cdc-monitoring/bin/pkg/elastic"
)

type MonitorRespository struct {
	es    *elasticsearch.Client
	index string
}

func NewElasticReposiroty() *MonitorRespository {
	return &MonitorRespository{
		es: elastic.ESCli,
	}
}

func (m *MonitorRespository) Index() error {
    return nil
}

func (m *MonitorRespository) Search() {
}

func (m *MonitorRespository) Update() {

}
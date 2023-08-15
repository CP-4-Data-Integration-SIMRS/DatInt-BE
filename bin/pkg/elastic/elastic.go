package elastic

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

var ESCli *elasticsearch.Client

func InitElastic() (err error) {

	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}

	ESCli, err = elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
		return
	}

	res, err := ESCli.Ping()

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Printf("ElasticSearch Connected %s",res)
	return
}

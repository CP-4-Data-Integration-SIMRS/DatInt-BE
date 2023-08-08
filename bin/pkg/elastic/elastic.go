package elastic

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

var ESCli *elasticsearch.Client

func InitElastic() {

	es, err := elasticsearch.NewDefaultClient()

	if err != nil {
		log.Fatal(err)
		return
	}

	res, err := es.Ping()

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(res)
}

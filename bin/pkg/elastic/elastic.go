package elastic

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

var ESCli *elasticsearch.Client

func InitElastic() (err error) {

	cfg := elasticsearch.Config{
		CloudID: "es-dbt:YXNpYS1zb3V0aGVhc3QxLmdjcC5lbGFzdGljLWNsb3VkLmNvbSRkMjYyNWJjNzY4NjA0ZDM1YTkzOWQyNWU2ZjI0NmJjMCQyMWI3Mjg3MjY2OWY0OTBmOTU3MTk1MjQ4ZGQ3YWNmNg==",
		APIKey:  "SDZ0Nko0b0JyTnVOd2FaWVN1NHI6WHhhZkNYQ1RSb1dtcU0zWUN4YUQxdw==",
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

	log.Printf("ElasticSearch Connected %s", res)
	return
}

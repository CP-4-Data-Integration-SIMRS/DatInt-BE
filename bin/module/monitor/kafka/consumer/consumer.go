package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/IBM/sarama"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/vier21/simrs-cdc-monitoring/bin/module/monitor/model"
	"github.com/vier21/simrs-cdc-monitoring/bin/pkg/elastic"
)

func main() {
	err := elastic.InitElastic()
	if err != nil {
		log.Fatal(err.Error())
	}

	topic := "mntr5"
	worker, err := connectConsumer([]string{"localhost:9092"})
	if err != nil {
		panic(err)
	}

	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	fmt.Println("Consumer started ")
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	msgCount := 0

	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				msgCount++
				fmt.Printf("Received message Count %d: | Topic(%s) | Message(%s) \n", msgCount, string(msg.Topic), string(msg.Value))
				err := indexOrUpdateDocuments(msg.Value)
				if err != nil {
					log.Println(err)
					continue
				}
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgCount, "messages")

	if err := worker.Close(); err != nil {
		panic(err)
	}

}

func connectConsumer(brokersUrl []string) (sarama.Consumer, error) {

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create new consumer
	conn, err := sarama.NewConsumer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func indexBulkData(msg []byte) error {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"}, // Elasticsearch server address
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	var data []model.DatabaseInfo
	var actions []string

	err = json.Unmarshal(msg, &data)
	if err != nil {
		log.Printf("error sending to elastic %s \n", err.Error())
		return err
	}

	for _, item := range data {
		action := `{ "index" : { "_index": "monitoring" } }`
		source, err := json.Marshal(item)
		if err != nil {
			return err
		}
		actions = append(actions, action, string(source))
	}

	body := strings.Join(actions, "\n") + "\n"

	res, err := client.Bulk(
		strings.NewReader(body),
		client.Bulk.WithContext(context.Background()),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("bulk indexing failed: %s", res.Status())
	}

	fmt.Println("Document indexed or updated successfully")

	return nil
}

func PushToElastic(message []byte) error {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"}, // Elasticsearch server address
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	var data model.DatabaseInfo

	err = json.Unmarshal(message, &data)
	if err != nil {
		log.Printf("error sending to elastic %s \n", err.Error())
		return err
	}

	docJSON, err := json.Marshal(data)
	if err != nil {
		log.Printf("error sending to elastic %s \n", err.Error())
		return err
	}

	req := esapi.IndexRequest{
		Index:      "monitoring",
		DocumentID: "1",
		Body:       strings.NewReader(string(docJSON)),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), client)

	if err != nil {
		log.Fatalf("Error indexing document: %s", err)
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error response: %s", res)
		return err
	}

	fmt.Println("Document indexed or updated successfully")
	return nil
}

func indexOrUpdateDocuments(msg []byte) error {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"}, // Elasticsearch server address
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	var items []model.DatabaseInfo
	err = json.Unmarshal(msg, &items)
	if err != nil {
		log.Printf("error sending to elastic %s \n", err.Error())
		return err
	}

	for _, item := range items {
		docID := item.ID

	

		itemJSON, err := json.Marshal(item)
		if err != nil {
			return err
		}

		req := esapi.IndexRequest{
			Index:      "mntr5",
			DocumentID: docID,
			Body:       strings.NewReader(string(itemJSON)),
			Refresh:    "true",
		}

		res, err := req.Do(context.Background(), client)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.IsError() {
			return fmt.Errorf("Error indexing/updating document: %s", res.Status())
		}
	}
	fmt.Println("Document indexed or updated successfully")

	return nil
}

func documentExists(client *elasticsearch.Client, docname string) (bool, error) {
	res, err := client.Exists("mntr", docname)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return false, nil
	} else if res.IsError() {
		return false, fmt.Errorf("Error checking document existence: %s", res.Status())
	}

	return true, nil
}

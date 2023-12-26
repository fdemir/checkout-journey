package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type Checkout struct {
	Address  string    `json:"address"`
	Products []Product `json:"products"`
}

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BROKER"),
		"group.id":          "shipment",
		"auto.offset.reset": "earliest",
		"fetch.min.bytes":   1,
	})

	if err != nil {
		panic(err)
	}
	topic := "checkout"

	err = c.SubscribeTopics([]string{topic}, nil)

	if err != nil {
		panic(err)
	}

	defer c.Close()

	// TODO: concurrent event processing
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {

			var checkout Checkout

			dec := json.NewDecoder(bytes.NewReader(msg.Value))
			dec.Decode(&checkout)

			fmt.Println("Shipment should be to ", checkout.Address)
		} else {
			println("Error: " + err.Error())
		}
	}

}

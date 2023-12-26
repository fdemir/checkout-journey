package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Checkout struct {
	Address string `json:"address"`
	Email   string `json:"email"`
}

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BROKER"),
		"group.id":          "notification",
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

			// TODO: implement email notification
			fmt.Println("Notification should send to ", checkout.Email)
		} else {
			println("Error: " + err.Error())
		}
	}

}

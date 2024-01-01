package main

import (
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BROKER"),
		"group.id":          "inventory",
		"auto.offset.reset": "earliest",
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
			println("Items sold:")
			println(string(msg.Value))
		} else {
			println("Error: " + err.Error())
		}
	}

}

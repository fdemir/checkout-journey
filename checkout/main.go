package main

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gofiber/fiber/v2"
)

const (
	topic = "checkout"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type Checkout struct {
	Address  string    `json:"address"`
	Email    string    `json:"email" validate:"required,email"`
	Products []Product `json:"products"`
}

func main() {
	app := fiber.New()

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BROKER"),
		"linger.ms":         0,
	})

	if err != nil {
		panic(err)
	}

	defer p.Close()

	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"name":   "checkout",
			"status": "ok",
		})
	})

	app.Post("/checkout", func(c *fiber.Ctx) error {
		var checkout Checkout

		if err := c.BodyParser(&checkout); err != nil {
			return err
		}

		targetTopic := topic

		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &targetTopic, Partition: kafka.PartitionAny},
			Value:          []byte(c.Body()),
		}, nil)

		return c.JSON(checkout)
	})

	log.Fatal(app.Listen(":3000"))
}

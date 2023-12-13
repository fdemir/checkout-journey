package main

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gofiber/fiber/v2"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

type Checkout struct {
	Products []Product `json:"products"`
}

func main() {
	app := fiber.New()

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BROKER"),
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

		topic := "checkout"

		// create topic if not exists

		for _, product := range checkout.Products {
			p.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
				Value:          []byte(product.Name),
			}, nil)
		}

		return c.JSON(checkout)
	})

	log.Fatal(app.Listen(":3000"))
}

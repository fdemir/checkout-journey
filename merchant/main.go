package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Merchant Service",
		})
	})

	app.Get("/products", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"products": []map[string]interface{}{
				{"id": 1001001, "name": "AromaBrew", "price": 15.99},
				{"id": 1001002, "name": "BeanEssence", "price": 18.50},
				{"id": 1001003, "name": "CaféDelight", "price": 20.00},
				{"id": 1001004, "name": "MochaMagic", "price": 17.75},
				{"id": 1001005, "name": "EspressoElixir", "price": 22.30},
				{"id": 1001006, "name": "MorningMist", "price": 16.50},
				{"id": 1001007, "name": "SunriseBlend", "price": 19.95},
				{"id": 1001008, "name": "RoastRevival", "price": 21.40},
				{"id": 1001009, "name": "BrewedBliss", "price": 14.99},
				{"id": 1001010, "name": "CaramelCraze", "price": 23.65},
			},
		})
	})

	log.Fatal(app.Listen(":3004"))
}

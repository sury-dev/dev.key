package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sury-dev/dev-key-backend/config"
)

func main() {

	config.LoadEnv()

	port := config.GetEnv("PORT", "8080")

	app := fiber.New(fiber.Config{
		AppName: "Dev.Key",
	})

	allowedOrigins := map[string]bool{
		"http://localhost:4200": true,
	}

	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return allowedOrigins[origin]
		},
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Dev.Key Server is running",
		})
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}

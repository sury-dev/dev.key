package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sury-dev/dev-key-backend/config"
	"github.com/sury-dev/dev-key-backend/database"
)

func main() {

	// Load the environment variables
	config.LoadEnv()

	// Get the port from the environment variables
	port := config.GetEnv("PORT", "8080")

	// Ping the MongoDB database
	//This will first connect to the MongoDB ( if not already connected, else it will use the already connected client ) database and then ping it
	if err := database.Ping(); err != nil {
		log.Fatal("Error pinging MongoDB", err)
	}

	// Create a new Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Dev.Key",
	})

	// Allowed origins
	// This is the allowed origins for the CORS policy
	allowedOrigins := map[string]bool{
		"http://localhost:4200": true,
	}

	// CORS policy
	app.Use(cors.New(cors.Config{
		AllowOriginsFunc: func(origin string) bool {
			return allowedOrigins[origin]
		},
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))

	// Health check endpoint
	// This is a simple health check endpoint to check if the server is running
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Dev.Key Server is running",
		})
	})

	// Start the server
	// This will start the server on the port specified in the environment variables
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}

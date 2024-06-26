package main

import (
	"fmt"
	"log"

	"github.com/aditansh/blendnet-task/be/config"
	"github.com/aditansh/blendnet-task/be/database"
	"github.com/aditansh/blendnet-task/be/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {

	app := fiber.New()

	config, err := config.LoadEnvVariables(".")
	if err != nil {
		log.Fatalf("Error loading env variables: %s", err.Error())
	}
	database.ConnectDB(&config)

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.ClientOrigin,
		AllowHeaders:     "*",
		AllowMethods:     "GET, POST, PUT, PATCH, DELETE, OPTIONS",
		AllowCredentials: true,
	}))

	fmt.Println("Server started")

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  true,
			"message": "pong",
		})
	})

	routes.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"status":  false,
			"message": "Route Not Found",
		})
	})

	log.Fatal(app.Listen(config.Port))
}

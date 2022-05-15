package main

import (
	"log"
	"todo_server/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.ConfigDefault))
	app.Use(logger.New(logger.Config{
		Format: "${status} - ${method} ${path}\n",
	}))

	app.Get("/", handlers.GetTodos)
	app.Get("/:id", handlers.GetTodo)
	app.Post("/", handlers.AddTodo)

	log.Fatal(app.Listen(":4000"))
}

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/m0hammedimran/todo_server/handlers"
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
	app.Put("/:id", handlers.UpdateTodo)

	log.Fatal(app.Listen(":4000"))
}

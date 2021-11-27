package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Completed bool
	Id        int
	Text      string
}

var todos = make([]Todo, 0)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		todo := append(todos, Todo{false, 1, "hello"})
		copy(todos[1:], todos[1+1:])

		// a[len(a)-1] = nil // or the zero value of T
		// a = a[:len(a)-1]

		return c.JSON(todo)
	})

	log.Fatal(app.Listen(":4000"))
}

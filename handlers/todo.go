package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Completed bool   `json:"completed"`
	Id        int    `json:"id"`
	Text      string `json:"text"`
}

var todos = []Todo{}

func GetTodos(c *fiber.Ctx) error {
	return c.JSON(todos)
}

func GetTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(fiber.Map{"error": "invalid id: " + c.Params("id")})
	}

	var todo Todo

	for _, t := range todos {
		if t.Id == id {
			todo = t
		}
	}

	if (todo == Todo{}) {
		return c.JSON(fiber.Map{"error": "todo not found"})
	}

	return c.JSON(todo)
}

func AddTodo(c *fiber.Ctx) error {
	var todo Todo

	if err := c.BodyParser(&todo); err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	todos = append(todos, todo)

	return c.JSON(fiber.Map{"message": "success"})
}

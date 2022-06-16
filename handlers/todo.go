package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/m0hammedimran/todo_server/lib"
	"github.com/m0hammedimran/todo_server/models"
)

// var todos = []Todo{}

func GetTodos(c *fiber.Ctx) error {
	var todos = []models.Todo{}
	lib.DB.Find(&todos)
	return c.JSON(todos)
}

func GetTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(fiber.Map{"error": "invalid id: " + c.Params("id")})
	}

	var todo models.Todo

	lib.DB.Find(&todo, id)

	return c.JSON(todo)
}

func AddTodo(c *fiber.Ctx) error {
	var todo models.Todo

	if err := c.BodyParser(&todo); err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	lib.DB.Create(&todo)

	return c.JSON(fiber.Map{"message": "success"})
}

func UpdateTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(fiber.Map{"error": "invalid id: " + c.Params("id")})
	}

	var todo models.Todo

	if err := c.BodyParser(&todo); err != nil {
		return c.JSON(fiber.Map{"error": err.Error()})
	}

	lib.DB.Model(&todo).Where("id = ?", id).Updates(&todo)

	return c.JSON(fiber.Map{"message": "success"})
}

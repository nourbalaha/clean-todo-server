package controllers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nourbalaha/clean-todo-server/internal/services"
	"github.com/nourbalaha/clean-todo-server/pkg/todo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoController struct {
	service *services.TodoService
}

func NewTodoController(service *services.TodoService) *TodoController {
	return &TodoController{
		service: service,
	}
}

func (c *TodoController) GetAllTodos(ctx echo.Context) error {
	todos, err := c.service.GetAllTodos()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, todos)
}

func (c *TodoController) GetTodoByID(ctx echo.Context) error {
	id := ctx.QueryParam("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid ID")
	}

	todo, err := c.service.GetTodoByID(objectID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	if todo == nil {
		return ctx.JSON(http.StatusNotFound, "Todo not found")
	}

	return ctx.JSON(http.StatusOK, todo)
}

func (c *TodoController) CreateTodo(ctx echo.Context) error {
	var todo todo.Todo
	if err := ctx.Bind(&todo); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	if err := c.service.CreateTodo(todo); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, "Todo created successfully")
}

func (c *TodoController) UpdateTodo(ctx echo.Context) error {
	var todo todo.Todo
	if err := ctx.Bind(&todo); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.service.UpdateTodo(todo); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "Todo updated successfully")
}

func (c *TodoController) DeleteTodo(ctx echo.Context) error {
	id := ctx.QueryParam("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid ID")
	}

	if err := c.service.DeleteTodo(objectID); err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusNoContent, nil)
}

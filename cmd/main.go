package main

import (
	"context"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/nourbalaha/clean-todo-server/internal/controllers"
	"github.com/nourbalaha/clean-todo-server/internal/repository"
	"github.com/nourbalaha/clean-todo-server/internal/services"
)

func main() {
	// MongoDB connection settings
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	database := client.Database("todoapp")
	todoCollection := database.Collection("todos")

	todoRepo := repository.NewMongoTodoRepository(todoCollection)
	todoService := services.NewTodoService(todoRepo)
	todoController := controllers.NewTodoController(todoService)

	// Initialize Echo framework
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/api/todos", todoController.GetAllTodos)
	e.GET("/api/todos/:id", todoController.GetTodoByID)
	e.POST("/api/todos", todoController.CreateTodo)
	e.PUT("/api/todos/:id", todoController.UpdateTodo)
	e.DELETE("/api/todos/:id", todoController.DeleteTodo)
	

	// Start the server
	e.Logger.Fatal(e.Start(":8000"))
}

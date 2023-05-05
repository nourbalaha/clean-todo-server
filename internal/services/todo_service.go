package services

import (
	"time"

	"github.com/nourbalaha/clean-todo-server/internal/repository"
	"github.com/nourbalaha/clean-todo-server/pkg/todo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoService struct {
	repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) *TodoService {
	return &TodoService{
		repo: repo,
	}
}

func (s *TodoService) GetAllTodos() ([]todo.Todo, error) {
	return s.repo.GetAllTodos()
}

func (s *TodoService) GetTodoByID(id primitive.ObjectID) (*todo.Todo, error) {
	return s.repo.GetTodoByID(id)
}

func (s *TodoService) CreateTodo(todo todo.Todo) error {
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	return s.repo.CreateTodo(todo)
}

func (s *TodoService) UpdateTodo(todo todo.Todo) error {
	todo.UpdatedAt = time.Now()
	return s.repo.UpdateTodo(todo)
}

func (s *TodoService) DeleteTodo(id primitive.ObjectID) error {
	return s.repo.DeleteTodo(id)
}

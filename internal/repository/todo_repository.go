package repository

import (
	"context"
	"time"

	"github.com/nourbalaha/clean-todo-server/pkg/todo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepository interface {
	GetAllTodos() ([]todo.Todo, error)
	GetTodoByID(id primitive.ObjectID) (*todo.Todo, error)
	CreateTodo(todo todo.Todo) error
	UpdateTodo(todo todo.Todo) error
	DeleteTodo(id primitive.ObjectID) error
}

type MongoTodoRepository struct {
	collection *mongo.Collection
}

func NewMongoTodoRepository(collection *mongo.Collection) *MongoTodoRepository {
	return &MongoTodoRepository{
		collection: collection,
	}
}

func (r *MongoTodoRepository) GetAllTodos() ([]todo.Todo, error) {
	var todos []todo.Todo

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

func (r *MongoTodoRepository) GetTodoByID(id primitive.ObjectID) (*todo.Todo, error) {
	var todo todo.Todo

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&todo)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // No todo found with the given ID
		}
		return nil, err
	}

	return &todo, nil
}

func (r *MongoTodoRepository) CreateTodo(todo todo.Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, todo)
	return err
}

func (r *MongoTodoRepository) UpdateTodo(todo todo.Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"title":       todo.Title,
			"description": todo.Description,
			"completed":   todo.Completed,
			"updatedAt":   time.Now(),
		},
	}

	_, err := r.collection.UpdateOne(ctx, bson.M{"_id": todo.ID}, update)
	return err
}

func (r *MongoTodoRepository) DeleteTodo(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
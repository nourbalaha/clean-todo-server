package todo

import "time"

type Todo struct {
	ID          string    `bson:"_id,omitempty"`
	Title       string    `bson:"title,omitempty"`
	Description string    `bson:"description,omitempty"`
	Completed   bool      `bson:"completed,omitempty"`
	CreatedAt   time.Time `bson:"createdAt,omitempty"`
	UpdatedAt   time.Time `bson:"updatedAt,omitempty"`
}
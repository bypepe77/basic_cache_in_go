package models

import "github.com/google/uuid"

type Todo struct {
	ID       uuid.UUID
	Content  string
	Finished bool
}

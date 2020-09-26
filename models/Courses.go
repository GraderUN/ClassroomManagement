package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Courses ...

type Courses struct {
	ID    primitive.ObjectID `json:"id"`
	grado int8               `json:"grado"`
	letra string             `json:"letra"`
}

package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AssignedClassroom ..
type AssignedClassroom struct {
	ID       primitive.ObjectID `json:"id"`
	course   primitive.ObjectID `json:"id"`
	letra    primitive.ObjectID `json:"letra"`
	profesor primitive.ObjectID `json:"profesor"`
}

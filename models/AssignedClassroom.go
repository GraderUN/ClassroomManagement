package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AssignedClassroom ..
type AssignedClassroom struct {
	Course   primitive.ObjectID `json:"id"`
	Letra    primitive.ObjectID `json:"letra"`
	Profesor int                `json:"profesor"`
}

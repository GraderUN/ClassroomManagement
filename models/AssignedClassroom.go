package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AssignedClassroom ..
type AssignedClassroom struct {
	Course   primitive.ObjectID `json:"curso"`
	Salon    primitive.ObjectID `json:"salon"`
	Profesor int                `json:"profesor"`
	Horario  string             `json:"horario"`
}

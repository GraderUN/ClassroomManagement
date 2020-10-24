package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// AssignedClassroom ..

type AssignedClassroom struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Curso    string             `json:"curso"`
	Salon    string             `json:"salon"`
	Profesor string             `json:"profesor"`
	Horario  string             `json:"horario"`
	Materia  string             `json:"materia"`
}

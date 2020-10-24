package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Courses ...

type Courses struct {
	Id            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Grado         int8               `json:"grado"`
	Letra         string             `json:"letra"`
	Id_estudiante []string           `json:"id_estudiante"`
}

package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Type ...

type Classroom struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Capacidad   int                `json:"capacidad"`
	Description string             `json:"description"`
}

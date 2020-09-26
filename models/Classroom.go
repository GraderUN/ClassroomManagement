package models

//Type ...

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Classroom struct {
	ID          primitive.ObjectID `json:"id"`
	description string             `json:"description"`
}

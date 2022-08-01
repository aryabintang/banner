package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Banner struct {
	Id     primitive.ObjectID `json:"id,omitempty"`
	Banner string             `json:"banner,omitempty" validate:"required"`
	Alt    string             `json:"alt,omitempty" validate:"required"`
	Link   string             `json:"link,omitempty" validate:"required"`
}

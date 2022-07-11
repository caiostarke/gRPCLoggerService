package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Log struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Message string             `bson:"message"`
}

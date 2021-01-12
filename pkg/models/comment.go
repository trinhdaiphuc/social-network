package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Body      string             `bson:"body" json:"body"`
	Username  string             `bson:"username" json:"username"`
	CreatedAt string             `bson:"createdAt" json:"createdAt"`
}

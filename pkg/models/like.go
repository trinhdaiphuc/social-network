package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Like struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username  string             `bson:"username" json:"username"`
	CreatedAt string             `bson:"createdAt" json:"createdAt"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Body      string             `bson:"body" json:"body"`
	CreatedAt string             `bson:"createdAt" json:"createdAt"`
	Username  string             `bson:"username" json:"username"`
	Comments  []Comment          `bson:"comments,omitempty" json:"comments"`
	Likes     []Like             `bson:"likes,omitempty" json:"likes"`
}

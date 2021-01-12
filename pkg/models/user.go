package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email     string             `bson:"email" json:"email"`
	Username  string             `bson:"username" json:"username"`
	Password  string             `bson:"password" json:"-"`
	Token     string             `bson:"-" json:"token"`
	CreatedAt string             `bson:"createdAt" json:"createdAt"`
}

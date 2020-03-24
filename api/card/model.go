package card

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type APICard struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	Name       string             `json:"name" bson:"name"`
	Content    string             `json:"content" bson:"content"`
	ModifyTime int64
	CreateTime int64
	CreateUser primitive.ObjectID
}

package user

import "go.mongodb.org/mongo-driver/bson/primitive"

// User 用户信息表
type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"mail" bson:"email"`
	Password string             `json:"password" bson:"password"`
}

package user

import (
	"eksapi/global/setting"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Key     []byte
	UserCol *mongo.Collection
)

func init() {
	Key = []byte(setting.Config.Secret)
}

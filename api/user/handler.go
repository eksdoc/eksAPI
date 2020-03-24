package user

import (
	"context"
	"eksapi/utils/statuscode"
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type _loginResq struct {
	Name     string `json:"name" bson:"name" valid:"required"`
	Password string `json:"password" bson:"password" valid:"required"`
}

/* post users/login

用户登录
*/
func Login(c *gin.Context) {
	var err error

	// 请求体
	body, err := c.GetRawData()
	if err != nil {
		c.Status(statuscode.BadRequest)
		return
	}
	// 解析与验证
	resq := _loginResq{}
	if err = json.Unmarshal(body, &resq); err != nil {
		c.Status(statuscode.BadRequest)
		return
	}
	if ok, _ := govalidator.ValidateStruct(resq); !ok {
		c.Status(statuscode.BadRequest)
		return
	}
	// 读取数据库
	user := User{}
	err = UserCol.FindOne(context.Background(), bson.M{"name": resq.Name, "password": resq.Password}).Decode(&user)
	if err != nil {
		c.Status(statuscode.BadRequest)
		return
	}
	// token
	token, err := GenerateToken(user.ID, user.Name)
	if err != nil {
		c.Status(statuscode.BadRequest)
		return
	}

	c.JSON(statuscode.OK, gin.H{"token": token})
}

type _registerResq struct {
	Name     string `json:"name" bson:"name" valid:"required"`
	Email    string `json:"name" bson:"name" valid:"email"`
	Password string `json:"password" bson:"password" valid:"required"`
}

/* post users/register

用户注册
*/
func Register(c *gin.Context) {
	var err error

	// 请求体
	body, err := c.GetRawData()
	if err != nil {
		c.Status(statuscode.BadRequest)
		return
	}
	// 解析与验证
	resq := _registerResq{}
	if err = json.Unmarshal(body, &resq); err != nil {
		c.Status(statuscode.BadRequest)
		return
	}
	if ok, _ := govalidator.ValidateStruct(resq); !ok {
		c.Status(statuscode.BadRequest)
		return
	}
	// 存储数据
	_, err = UserCol.InsertOne(context.Background(), resq)
	if err != nil {
		c.Status(statuscode.InternalServerError)
		return
	}

	c.Status(statuscode.Created)
}

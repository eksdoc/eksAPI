package user

import (
	"eksapi/utils/statuscode"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// 自定义保存在token中的数据
type MyCustomClaims struct {
	ID   primitive.ObjectID `json:"id"`
	Name string             `json:"name"`
	jwt.StandardClaims
}

// 生成token
func GenerateToken(id primitive.ObjectID, name string) (string, error) {
	// 设置token
	claims := MyCustomClaims{
		id,
		name,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Second * 60 * 60 * 24 * 7).Unix(),
		},
	}

	// 按照指定的算法生成token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return token.SignedString(Key)
}

// 验证token
func ParseToken(tokenString string) (*MyCustomClaims, bool) {
	time.Unix(0, 0)
	token, _ := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Key, nil
	})

	// 如果时间过期也算验证失败
	if data, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return data, true
	} else {
		return data, false
	}
}

// 验证是否登录的中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头中的token
		token := c.GetHeader("eksToken")
		if claim, ok := ParseToken(token); ok {
			c.Set("userId", claim.ID)
			c.Set("userName", claim.Name)
			c.Next()
		} else {
			c.AbortWithStatusJSON(statuscode.Unauthorized, gin.H{"err": "未登录"})
		}
	}
}

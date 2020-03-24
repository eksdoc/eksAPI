package main

import (
	"eksapi/api/user"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	users := r.Group("users")
	user.Router(users)

}

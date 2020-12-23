package main

import (
	"gin-redis/src/controller"
	"gin-redis/src/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))

	r.Use(sessions.Sessions("mysession", store))

	r.POST("/login", controller.Login)
	r.GET("/logout", controller.Logout)

	auth := r.Group("/auth")
	auth.Use(middleware.Authentication())

	auth.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	r.Run()
}

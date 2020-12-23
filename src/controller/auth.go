package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("id", 12121212)
	session.Set("email", "test@test.com")
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"message": "User sign in successfully",
	})
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"message": "User sign out successfully",
	})
}

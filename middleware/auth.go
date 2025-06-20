package middleware

import (
	"koreader-sync-server-go/db"
	"koreader-sync-server-go/models"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	user := c.GetHeader("x-auth-user")
	key := c.GetHeader("x-auth-key")
	var u models.User
	if err := db.DB.Where("username = ? AND key_md5 = ?", user, key).First(&u).Error; err != nil {
		c.JSON(401, gin.H{"code": 2001, "message": "Unauthorized"})
		c.Abort()
		return
	}
	c.Set("user_id", u.ID)
	c.Set("username", u.Username)
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Next()
}

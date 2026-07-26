package middlewares

import (
	"go-project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized!",
		})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "not authorized!",
		})
		return
	}

	//like intent
	c.Set("userId", userId)
	c.Next()

	//بعد از اتمام درسخئاست

}

package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func VerifyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionTokens := session.Get("session_tokens")

		if sessionTokens == nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Unauthorized"})
			return
		}

		sessions, exists := c.Get("sessions")

		if sessions == nil || !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		c.Next()
	}
}

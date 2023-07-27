package middleware

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	m "smapurv1_api/models"
	s "smapurv1_api/setup"
)

func DeserializeUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := VerifyAuth(); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err})
			return
		}

		session := sessions.Default(c)
		getSession := session.Get("session_tokens")

		id := reflect.ValueOf(getSession).Elem().Field(0).Interface()

		var user m.Users
		result := s.DB.First(&user, "kd_user = ?", fmt.Sprint(id))
		if result.Error != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "user belonging to this session no longer exists"})
			return
		}

		c.Set("currentUser", user)
		c.Next()
	}
}

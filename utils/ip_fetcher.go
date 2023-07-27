package utils

import (
	"net"
	"strings"

	"github.com/gin-gonic/gin"
)

func FetchUserIP(c *gin.Context) string {
	var userIp string

	if len(c.Request.Header.Get("CF-Connecting-IP")) > 1 {
		userIp = c.Request.Header.Get("CF-Connecting-IP")
		return string(net.ParseIP(userIp))
	} else if len(c.Request.Header.Get("X-Forwarded-For")) > 1 {
		userIp = c.Request.Header.Get("X-Forwarded-For")
		return string(net.ParseIP(userIp))
	} else if len(c.Request.Header.Get("X-Real-IP")) > 1 {
		userIp = c.Request.Header.Get("X-Real-IP")
		return string(net.ParseIP(userIp))
	} else {
		userIp = c.Request.RemoteAddr
		if strings.Contains(userIp, ":") {
			return string(net.ParseIP(strings.Split(userIp, ":")[0]))
		} else {
			return string(net.ParseIP(userIp))
		}
	}
}

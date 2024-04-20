package network

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (n *Network) verifyLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := getAuthToken(c)
		if t == "" {
			c.JSON(http.StatusUnauthorized, nil)
			c.Abort()
		} else {
			if _, err := n.gRPCClient.VerifyAuth(t); err != nil {
				c.JSON(http.StatusUnauthorized, err)
				c.Abort()
			} else {
				c.Next()
			}
		}
	}
}

func getAuthToken(c *gin.Context) string {
	var token string

	authToken := c.Request.Header.Get("Authorization")
	authSlided := strings.Split(authToken, " ")

	if len(authSlided) > 1 {
		token = authSlided[1]
	}

	return token
}

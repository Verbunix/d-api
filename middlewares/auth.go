package middlewares

import (
	"d-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JwtTokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := ExtractJwtToken(c.Request)
		if err := services.TokenValidation(token); err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}

		c.Next()
	}
}

func ExtractJwtToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

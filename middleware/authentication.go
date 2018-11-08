package middleware

import (
	"gin_todo/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("token")

		claims := util.ParseToken(tokenString)

		if claims != nil {
			c.Set("UserId", int64(claims["id"].(float64)))
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "unauthered",
				"data": "",
			})

			c.Abort()
			return
		}

	}
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func VisitorCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		visitor, err := c.Cookie("visitor")

		if err != nil || visitor == "" {
			visitor = uuid.New().String()

			c.SetCookie(
				"visitor",
				visitor,
				60*60*24*30,
				"/",
				"",
				false,
				true,
			)
		}

		c.Next()
	}
}
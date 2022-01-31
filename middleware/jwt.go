package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func IsAuthorized(endpoint func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			claim := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(c.GetHeader("Authorization"), claim, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}

				return []byte("LP2KXyA6ZZds3JDRY9XuGbUDrPWndPPq"), nil
			})
			if err != nil {
				fmt.Fprintf(c.Writer, err.Error())
			}
			if token.Valid {
				c.Set("token", claim["sub"])
				endpoint(c)
			}
		} else {
			fmt.Fprintf(c.Writer, "Not Authorized")
		}
	}
}

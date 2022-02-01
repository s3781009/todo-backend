package middleware

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func IsAuthorized(endpoint func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			claim := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(c.GetHeader("Authorization"), claim, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}

				read, err := godotenv.Read(".env")
				if err != nil {
					log.Fatal("cannot read env")
				}
				return []byte(read["AUTH0_CLIENT_SECRET"]), nil
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

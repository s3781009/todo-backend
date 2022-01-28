package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var Claims jwt.MapClaims

func IsAuthorized(endpoint func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			claim := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(c.GetHeader("Authorization"), claim, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}
				err := godotenv.Load(".env")
				if err != nil {
					panic(err.Error())
				}
				read, err := godotenv.Read(".env")
				if err != nil {
					panic(err.Error())
				}
				return []byte(read["AUTH0_CLIENT_SECRET"]), nil
			})
			if err != nil {
				fmt.Fprintf(c.Writer, err.Error())
			}
			if token.Valid {
				Claims = claim
				fmt.Println(Claims["sub"])
				endpoint(c)
			}
		} else {
			fmt.Fprintf(c.Writer, "Not Authorized")
		}
	}
}
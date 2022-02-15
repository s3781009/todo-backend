package middleware

import (
	"fmt"

	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// IsAuthorized middleware for authenticating user with auth0 , validating request access token against secret
//key, the route endpoint is  passed in as a callback and is only called if the access token is validate
func IsAuthorized(endpoint func(c *gin.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			claim := jwt.MapClaims{}
			token, err := jwt.ParseWithClaims(c.GetHeader("Authorization"), claim, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}

				return []byte(os.Getenv("AUTH0_CLIENT_SECRET")), nil
			})
			if err != nil {
				fmt.Fprintf(c.Writer, err.Error())
			}
			if token.Valid {
				c.Set("token", claim["sub"]) //set the 'sub' claim from access token as the 'user-id'
				endpoint(c)
			}
		} else {
			fmt.Fprintf(c.Writer, "Not Authorized")
		}
	}
}

package middleware

import (
	"net/http"

	"gin/config"
	"gin/helper"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Mengambil cookie token
		cookie, err := c.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// Jika cookie tidak ada, return unauthorized
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJson(c.Writer, http.StatusUnauthorized, response)
				c.Abort()
				return
			}
		}

		// Mengambil token value
		tokenString := cookie

		claims := &config.JWTClaim{}
		// Parsing token jwt
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				// Jika token invalid
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJson(c.Writer, http.StatusUnauthorized, response)
				c.Abort()
				return
			case jwt.ValidationErrorExpired:
				// Jika token expired
				response := map[string]string{"message": "Unauthorized, Token expired!"}
				helper.ResponseJson(c.Writer, http.StatusUnauthorized, response)
				c.Abort()
				return
			default:
				// Jika error lainnya
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJson(c.Writer, http.StatusUnauthorized, response)
				c.Abort()
				return
			}
		}

		if !token.Valid {
			// Jika token tidak valid
			response := map[string]string{"message": "Unauthorized"}
			helper.ResponseJson(c.Writer, http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		// Lanjutkan ke handler berikutnya
		c.Next()
	}
}

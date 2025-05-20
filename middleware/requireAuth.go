package middleware

import (
	"net/http"
	"os"
	"time"

	"github.com/Divyshekhar/go-jwt-api/intializers"
	"github.com/Divyshekhar/go-jwt-api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(ctx *gin.Context) {
	//get the cookie
	tokenString, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	//decode and validate

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp, ok := claims["exp"].(float64)         //jwt payload stored in float64 so parsing it safely
		if !ok || int64(exp) < time.Now().Unix() { //time.Now().Unix() is an int64 so convert exp which was float64
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var user models.User
		intializers.DB.First(&user, claims["sub"])

		if user.ID == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("user", user)
		ctx.Next()

	} else {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

}

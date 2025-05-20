package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/Divyshekhar/go-jwt-api/intializers"
	"github.com/Divyshekhar/go-jwt-api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(ctx *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to read body"})
		return
	}
	//hash the password
	hash, er := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if er != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to hash password"})
		return
	}
	user := models.User{Email: body.Email, Password: string(hash)}
	result := intializers.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to create user"})
		return
	}

	ctx.JSON(200, gin.H{"message": "User created successfully"})

}

func Login(ctx *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	ctx.ShouldBindJSON(&body)
	var user models.User
	intializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		ctx.JSON(400, gin.H{"error": "Invalid email or password"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid email or password"})
		return

	}
	//generate jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, eror := token.SignedString([]byte(os.Getenv("SECRET")))
	if eror != nil {
		ctx.JSON(400, gin.H{
			"Error": "failed to generate token",
		})
		return
	}
	
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)

	ctx.JSON(200, gin.H{})

}

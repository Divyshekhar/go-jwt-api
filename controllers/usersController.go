package controllers

import (
	"net/http"

	"github.com/Divyshekhar/go-jwt-api/intializers"
	"github.com/Divyshekhar/go-jwt-api/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(ctx *gin.Context) {
	var body struct {
		Email    string
		Password string
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

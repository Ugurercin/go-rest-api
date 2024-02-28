package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var u models.User
	err := ctx.ShouldBindJSON(&u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = u.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, u)
}

func login(ctx *gin.Context) {
	var u models.User

	err := ctx.ShouldBindJSON(&u)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = u.ValidateCrendentials()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "could not authenticate user"})
		return
	}

	token, err := utils.GenerateToken(u.Email, u.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not authenticate the user."})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "login successful", "token": token})

}

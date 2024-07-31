package routes

import (
	"event-booking-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindBodyWithJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data!"})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User created successfully!"})

}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindBodyWithJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse request data."})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login successful!"})
}

package routes

import (
	"go-project/models"
	"go-project/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createUser(c *gin.Context) {

	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "there was an error in data " + err.Error(),
		})
		return
	}

	err = user.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "there was an error On saving Data! Try Again!",
		})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{
		"status": "saved",
		"data":   user,
	})

}

func loginUser(c *gin.Context) {

	var user models.User

	err := c.ShouldBindJSON(&user)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "there was an error in data " + err.Error(),
		})
		return
	}
	err = user.ValidateUserLogin()

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "there was an error in data " + err.Error(),
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": "Could not authenticate token!",
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "the user is logged in",
		"token":   token,
	})

}

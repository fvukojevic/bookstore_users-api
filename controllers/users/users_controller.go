package users

import (
	"github.com/fvukojevic/bookstore_users-api/domain/users"
	"github.com/fvukojevic/bookstore_users-api/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		//TODO: Return bad request to the caller
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: Handle user creation error
		return
	}

	c.JSON(http.StatusCreated, result)
}

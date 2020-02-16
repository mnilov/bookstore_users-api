package users

import (
	"github.com/gin-gonic/gin"
	"github.com/mnilov/bookstore_users-api/domain/users"
	"github.com/mnilov/bookstore_users-api/services"
	"github.com/mnilov/bookstore_users-api/utils/errors"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json error")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveError := services.CreateUser(user)
	if saveError != nil {
		c.JSON(saveError.Status, saveError)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

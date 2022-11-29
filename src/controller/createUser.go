package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lucasr0drigues/GoCrud/src/configuration/rest_err"
	"github.com/lucasr0drigues/GoCrud/src/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorret fields, error=%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}

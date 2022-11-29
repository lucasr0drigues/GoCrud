package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/lucasr0drigues/GoCrud/src/configuration/rest_err"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Você chamou a rtoa de forma errada")
	c.JSON(err.Code, err)
}

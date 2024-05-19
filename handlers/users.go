package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	data := map[string]string{
		"message": "hello world",
	}
	c.JSON(http.StatusOK, data)
}

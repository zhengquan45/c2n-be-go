package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HelloController returns a greeting message
func GetHelloWord(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello, World!")
}

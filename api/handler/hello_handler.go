package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloHandler struct {
}

func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

// HelloController returns a greeting message
func (h *HelloHandler) GetHelloWord(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello, World!")
}

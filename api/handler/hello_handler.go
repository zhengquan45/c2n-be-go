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

// @Summary HelloWorld
// @Description Returns a greeting message
// @Tags hello
// @Accept json
// @Produce json
// @Success 200 {string} string "Hello, World!"
// @Router /hello [get]
func (h *HelloHandler) GetHelloWord(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello, World!")
}

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

// @Summary Get Hello World
// @Description Returns a greeting message
// @Tags hello
// @Accept json
// @Produce json
// @Success 200 {string} string "Hello, World!"
// @Failure 400 {string} string "Bad Request"
// @Failure 404 {string} string "Not Found"
// @Router /hello [get]

func (h *HelloHandler) GetHelloWord(c *gin.Context) {
	c.JSON(http.StatusOK, "Hello, World!")
}

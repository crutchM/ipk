package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//класс, для более простой генерации ошибки, и отправки ее клиенту
type Error struct {
	Message string `json:"message"`
}

func NewError(message string) *Error {
	return &Error{Message: message}
}
func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, NewError(message))
}

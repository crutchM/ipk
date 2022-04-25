package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"ipk/pkg/data"
	"net/http"
)

func (h *Handler) CreateTest(c *gin.Context) {
	var test data.Test
	if err := c.BindJSON(test); err != nil {
		logrus.Error(err.Error())
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.CreateTest(test)
	if err != nil {
		logrus.Error(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	SendJSONResponse(c, "testId", id)
}

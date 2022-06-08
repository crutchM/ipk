package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"ipk/domain/model"
	"net/http"
)

func (h *Handler) getAllChairs(c *gin.Context) {
	chairs, err := h.services.GetAllChairs()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	h.addHeaders(c)
	c.JSON(http.StatusOK, map[string]interface{}{
		"chair": chairs,
	})
}

func (h *Handler) createChair(c *gin.Context) {
	var input model.Chair
	if err := c.BindJSON(&input); err != nil {
		logrus.Error(err.Error())
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	chair, err := h.services.CreateChair(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	h.addHeaders(c)
	c.JSON(http.StatusOK, map[string]interface{}{
		"chair": chair,
	})
}

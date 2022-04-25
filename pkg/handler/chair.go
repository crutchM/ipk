package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"ipk"
	"net/http"
)

func (h *Handler) getAllChairs(c *gin.Context) {
	chairs, err := h.services.GetAllChairs()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"chair": chairs,
	})
}

func (h *Handler) createChair(c *gin.Context) {
	var input ipk.Chair

	if err := c.BindJSON(&input); err != nil {
		logrus.Error(err.Error())
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	chair, err := h.services.CreateChair(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"chair": chair,
	})
}

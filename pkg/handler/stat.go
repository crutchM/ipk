package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type InputStruct struct {
	Chair int `json:"chair"`
}

func (h *Handler) getStat(c *gin.Context) {
	var input InputStruct
	if err := c.BindJSON(&input); err != nil {
		logrus.Error(err.Error())
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	res, err := h.services.StatInterface.GetStat(input.Chair)
	if err != nil {
		logrus.Error(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h *Handler) getStatByTeacher(c *gin.Context) {
	id := c.Param("id")

	stat, err := h.services.GetStatByTeacher(id)
	if err != nil {
		logrus.Error(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, stat)

}

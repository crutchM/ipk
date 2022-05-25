package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) getStat(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	h.addHeaders(c)
	res, err := h.services.StatInterface.GetStat(id)
	if err != nil {
		logrus.Error(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)

}

func (h *Handler) getStatByTeacher(c *gin.Context) {
	id := c.Param("id")
	h.addHeaders(c)
	stat, err := h.services.GetStatByTeacher(id)
	if err != nil {
		logrus.Error(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, stat)

}

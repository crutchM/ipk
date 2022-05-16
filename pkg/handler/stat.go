package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"ipk/pkg/data/stat"
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

func (h *Handler) setStat(c *gin.Context) {
	var input stat.Stat
	if err := c.BindJSON(&input); err != nil {
		logrus.Error(err.Error())
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) setTestResult(c *gin.Context) {

}

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"ipk/pkg/data"
	"ipk/pkg/data/stat"
	"net/http"
	"strconv"
)

type Input struct {
	Name   string       `json:"name"`
	Blocks []data.Block `json:"blocks"`
}

func (h *Handler) CreateTest(c *gin.Context) {
	var input Input
	if err := c.BindJSON(&input); err != nil { //получаем данные из тела
		logrus.Error(err.Error())
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	test := data.Test{Name: input.Name, Blocks: input.Blocks}
	id, err := h.services.CreateTest(test) //записываем данные в базу
	if err != nil {
		logrus.Error(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	SendJSONResponse(c, "testId", id)
}

func (h *Handler) GetTest(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logrus.Error(err.Error())
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	test, er := h.services.GetTest(id)
	if err != nil {
		logrus.Error(er.Error())
		newErrorResponse(c, http.StatusInternalServerError, er.Error())
		return
	}
	c.JSON(http.StatusOK, test)
}

type ResInput struct {
	Blocks []data.Block `json:"blocks"`
	Test   int          `json:"test"`
}

func (h *Handler) SendResult(c *gin.Context) {
	var input ResInput
	if err := c.BindJSON(&input); err != nil {
		logrus.Error(err.Error())
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.AddResult(input.Blocks, input.Test)
	if err != nil {
		logrus.Error(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	SendJSONResponse(c, "status", "success")
}

func (h *Handler) SendStat(c *gin.Context) {
	var input stat.Stat
	if err := c.BindJSON(&input); err != nil {
		logrus.Error(err.Error())
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.AddRow(input)
	if err != nil {
		logrus.Error(err.Error())
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	SendJSONResponse(c, "rowId", id)
}

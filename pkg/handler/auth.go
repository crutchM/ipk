package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"ipk/pkg/data"
	"net/http"
)

// @SummarySignUp
// @Tags auth
// @ID create-account
// @Accept json
// @Produce json
// @Param input body data.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input data.User

	if err := c.BindJSON(&input); err != nil {
		logrus.Error(err.Error())
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorisation.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.Header("Access-Control-Allow-Origin", "http://192.168.11.40:3000")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "x-requested-with, Content-Type, origin, authorization, accept, x-access-token")
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
func (h *Handler) opt(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://192.168.11.40:3000")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "authorization, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	c.JSON(http.StatusOK, map[string]string{
		"cool": "cool",
	})
}

func (h *Handler) addHeaders(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://192.168.11.40:3000")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
}

func (h *Handler) check(c *gin.Context) {
	h.addHeaders(c)
	c.JSON(http.StatusOK, map[string]interface{}{
		"username": "alex",
		"password": "123",
	})
}

type signInInput struct {
	Username string `json:"username" `
	Password string `json:"password" `
}

func (h *Handler) signIn(c *gin.Context) {

	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		logrus.Error(err.Error())
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.services.Authorisation.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.Header("Access-Control-Allow-Origin", "http://192.168.11.40:3000")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	fmt.Println(token)
	SendJSONResponse(c, "token", token)
}

func (h *Handler) getAll(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://192.168.11.40:3000")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
	SendJSONResponse(c, "users", h.services.GetAll())
}

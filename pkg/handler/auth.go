package handler

import (
	"time"
	"net/http"
	"os"
	// "strings"

	"github.com/gin-gonic/gin"
	app "github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
)

func (h *Handler) InitAdmin(c *gin.Context) {
	input := app.User{
		Login:    os.Getenv("ADMIN_LOGIN"),
		Password: os.Getenv("ADMIN_PASSWORD"),
		Age:      os.Getenv("ADMIN_AGE"),
	}

	id, err := h.services.Authorization.CreateAdmin(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signUp(c *gin.Context) {
	var input app.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type userDataInput struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password_hash" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input userDataInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Login, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	red := app.GetRedisConn()
	_, err = red.Set(c, token.AccessUuid, token.AccessToken, 60*time.Second).Result()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token.AccessToken,
	})
}

func (h *Handler) logout(c *gin.Context) {
	claims, err := h.parseAuthHeader(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}

	red := app.GetRedisConn()
	_, err = red.Del(c, claims.AtUuid).Result()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, "Logout Success")
}

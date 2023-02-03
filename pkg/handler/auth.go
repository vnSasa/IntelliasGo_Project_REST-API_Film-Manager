package handler

import (
	"time"
	"net/http"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	app "github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/model"
)

func (h *Handler) InitAdmin(c *gin.Context) {
	input := app.User{
		Login:    viper.GetString("admin.login"),
		Password: viper.GetString("admin.password"),
		Age:      viper.GetString("admin.age"),
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
	at := time.Unix(token.AtExpires, 0)
	rt := time.Unix(token.RtExpires, 0)
	now := time.Now()
	_, err = red.Set(c, token.AccessUuid, token.AccessToken, at.Sub(now)).Result()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	_, err = red.Set(c, token.RefreshUuid, token.RefreshToken, rt.Sub(now)).Result()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"accsess_token": token.AccessToken,
		"refresh_token": token.RefreshToken,
	})
}

type refreshDataInput struct {
	RefreshToken string `json:"refresh_token"`
}

func (h *Handler) refreshSignIn(c *gin.Context) {
	var input refreshDataInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	RtData, err := h.services.Authorization.ParseRefreshToken(input.RefreshToken)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	if RtData.IsRefresh != true {
		newErrorResponse(c, http.StatusInternalServerError, "is not refresh token")

		return
	}

	token, err := h.services.Authorization.RefreshToken(RtData)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	
	red := app.GetRedisConn()
	_, err = red.Get(c, RtData.RtUuid).Result()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	_, err = red.Del(c, RtData.RtUuid).Result()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	_, err = red.Del(c, RtData.AtUuid).Result()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	
	at := time.Unix(token.AtExpires, 0)
	rt := time.Unix(token.RtExpires, 0)
	now := time.Now()
	_, err = red.Set(c, token.AccessUuid, token.AccessToken, at.Sub(now)).Result()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	_, err = red.Set(c, token.RefreshUuid, token.RefreshToken, rt.Sub(now)).Result()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"access_token": token.AccessToken,
		"refresh_token": token.RefreshToken,
	})
}

type logoutDataInput struct {
	AccessToken string `json:"access_token"`
}

func (h *Handler) logout(c *gin.Context) {
	var input logoutDataInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	claims, err := h.services.Authorization.ParseToken(input.AccessToken)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	red := app.GetRedisConn()
	_, err = red.Del(c, claims.AtUuid).Result()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}
	_, err = red.Del(c, claims.RtUuid).Result()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, "Logout Success")
}

package handler

import (
	"time"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	app "github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/model"
)

// @Summary InitAdmin
// @Tags auth
// @Description create admin
// @ID create admin
// @Accept json
// @Produce json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /init-admin [post]
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

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create account
// @Accept json
// @Produce json
// @Param input body app.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
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

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept json
// @Produce json
// @Param input body userDataInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
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

// @Summary RefreshSignIn
// @Tags auth
// @Description token
// @ID token
// @Accept json
// @Produce json
// @Param input body refreshDataInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/refresh [post]
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

// @Summary Logout
// @Tags auth
// @Description logout token
// @ID logout token
// @Accept json
// @Produce json
// @Param input body logoutDataInput true "credentials"
// @Success 200 {string} string "Logout Success"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/logout [post]
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

package handler

import (
	"errors"
	"net/http"
	"strings"

	app "github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/model"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) adminIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty auth header")

		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusBadRequest, "invalid auth header")

		return
	}
	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusBadRequest, "token is empty")

		return
	}
	claims, err := h.services.Authorization.VerifyAdminToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())

		return
	}
	red := app.GetRedisConn()
	_, err = red.Get(c, claims.AtUUID).Result()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.Set(userCtx, claims.UserID)
}

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusBadRequest, "empty auth header")

		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusBadRequest, "invalid auth header")

		return
	}
	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusBadRequest, "token is empty")

		return
	}
	claims, err := h.services.Authorization.VerifyUserToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "invalid token")

		return
	}
	red := app.GetRedisConn()
	_, err = red.Get(c, claims.AtUUID).Result()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "redis error")

		return
	}

	c.Set(userCtx, claims.UserID)
}

func getUserID(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "user id not found")

		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusBadRequest, "user id is invalid type")

		return 0, errors.New("user id not found")
	}

	return idInt, nil
}

package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/model"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) adminIdentity(c *gin.Context) {
	claims, err := h.parseAuthHeader(c)
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
	if !claims.IsAdmin {
		newErrorResponse(c, http.StatusUnauthorized, "only admin have access")

		return
	}
	c.Set(userCtx, claims.UserID)
}

func (h *Handler) userIdentity(c *gin.Context) {
	claims, err := h.parseAuthHeader(c)
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
	if !claims.IsUser {
		newErrorResponse(c, http.StatusUnauthorized, "only users have access")

		return
	}
	c.Set(userCtx, claims.UserID)
}

func (h *Handler) parseAuthHeader(c *gin.Context) (*app.AccessTokenClaims, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return nil, errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return nil, errors.New("invalid auth header")
	}

	if headerParts[0] != "Bearer" {
		return nil, errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return nil, errors.New("token is empty")
	}

	return h.services.Authorization.ParseToken(headerParts[1])
}

func getUserID(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")

		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is invalid type")

		return 0, errors.New("user id not found")
	}

	return idInt, nil
}

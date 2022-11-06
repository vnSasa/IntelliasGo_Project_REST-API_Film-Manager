package handler

import (
	"strings"
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

const (
	authorizationHeader = "Authorization"
	userCtx = "userId"
)

func (h *Handler) adminIdentity(c *gin.Context) {
	userId, err := h.parseAuthHeader(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	if userId != 1 {
		newErrorResponse(c, http.StatusUnauthorized, "only admin users have access")
		return
	}
	c.Set(userCtx, userId)
}

func (h *Handler) userIdentity(c *gin.Context) {
	userId, err := h.parseAuthHeader(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return 
	}
	c.Set(userCtx, userId)
}

func (h *Handler) parseAuthHeader(c *gin.Context) (int, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return 0, errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return 0, errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return 0, errors.New("token is empty")
	}

	return h.services.Authorization.ParseToken(headerParts[1])
}

func getUserId(c *gin.Context) (int, error) {
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
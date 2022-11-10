package handler

import (
	"os"
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
	userId, adminLogin, err := h.parseAuthHeader(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	if strings.Compare(adminLogin, "admin") != 0 {
		newErrorResponse(c, http.StatusUnauthorized, "only admin have access")
		return
	}
	c.Set(userCtx, userId)
}

func (h *Handler) userIdentity(c *gin.Context) {
	userId, userLogin, err := h.parseAuthHeader(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return 
	}
	if strings.Compare(userLogin, os.Getenv("ADMIN_LOGIN")) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "only users have access")
		return
	}
	c.Set(userCtx, userId)
}

func (h *Handler) parseAuthHeader(c *gin.Context) (int, string, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return 0, "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return 0, "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return 0, "", errors.New("token is empty")
	}

	return h.services.Authorization.ParseToken(headerParts[1])
}

func (h *Handler) getUserLoginById(c *gin.Context) (string, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return "", errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is invalid type")
		return "", errors.New("user id not found")
	}

	user, err := h.services.Authorization.GetUserById(idInt)
	if err != nil {
		return "", err
	}

	return user.Login, nil
}
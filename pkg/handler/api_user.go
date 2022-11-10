package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getDirectors(c *gin.Context) {
	user, err := h.getUserById(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"login": user.Login,
		"age": user.Age,
	})
}

func (h *Handler) getFilms(c *gin.Context) {
	user, err := h.getUserById(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"login": user.Login,
		"age": user.Age,
	})
}

func (h *Handler) createFavourite(c *gin.Context) {
	
}

func (h *Handler) getFavourite(c *gin.Context) {
	
}

func (h *Handler) getFavouriteById(c *gin.Context) {
	
}

func (h *Handler) updateFavourite(c *gin.Context) {
	
}

func (h *Handler) deleteFavourite(c *gin.Context) {
	
}

func (h *Handler) createWish(c *gin.Context) {
	
}

func (h *Handler) getWish(c *gin.Context) {
	
}

func (h *Handler) getWishById(c *gin.Context) {
	
}

func (h *Handler) updateWish(c *gin.Context) {
	
}

func (h *Handler) deleteWish(c *gin.Context) {
	
}
package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getDirectors(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getFilms(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
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
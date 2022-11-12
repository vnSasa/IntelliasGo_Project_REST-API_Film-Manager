package handler

import (
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createDiretor(c *gin.Context) {
	user, err := h.getUserById(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	var input app.DirectorList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.DirectorList.Create(user.Login, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllDiretors(c *gin.Context) {
	
}

func (h *Handler) getDiretorById(c *gin.Context) {
	
}

func (h *Handler) updateDiretor(c *gin.Context) {
	
}

func (h *Handler) deleteDiretor(c *gin.Context) {
	
}

func (h *Handler) createFilm(c *gin.Context) {
	
}

func (h *Handler) getAllFilms(c *gin.Context) {
	
}

func (h *Handler) getFilmById(c *gin.Context) {
	
}

func (h *Handler) updateFilm(c *gin.Context) {
	
}

func (h *Handler) deleteFilm(c *gin.Context) {
	
}
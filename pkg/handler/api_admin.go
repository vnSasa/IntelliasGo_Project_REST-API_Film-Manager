package handler

import (
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (h *Handler) createDiretor(c *gin.Context) {
	_, err := h.getUserById(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	
	var input app.DirectorList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.DirectorList.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllDirectorsResponce struct {
	Data []app.DirectorList `json:"data"`
}

func (h *Handler) getAllDiretors(c *gin.Context) {
	directors, err := h.services.DirectorList.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllDirectorsResponce{
		Data: directors,
	})
}

func (h *Handler) getDiretorById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	director, err := h.services.DirectorList.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, director)
}

func (h *Handler) updateDiretor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input app.UpdateDirectorInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.DirectorList.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteDiretor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	if err = h.services.DirectorList.Delete(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
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
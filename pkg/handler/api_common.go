package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
	"net/http"
	"strconv"
)

// GET ALL DIRECTORS...
type getAllDirectorsResponce struct {
	Directors []app.DirectorsList `json:"directors"`
}

func (h *Handler) getAllDiretors(c *gin.Context) {
	directors, err := h.services.DirectorsList.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, getAllDirectorsResponce{
		Directors: directors,
	})
}

// GET DIRECTOR BY ID...
func (h *Handler) getDiretorByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")

		return
	}

	director, err := h.services.DirectorsList.GetByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, director)
}

// GET ALL FILMS...
type getAllFilmsResponce struct {
	Films []app.FilmsList `json:"films"`
}

func (h *Handler) getAllFilms(c *gin.Context) {
	films, err := h.services.FilmsList.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, getAllFilmsResponce{
		Films: films,
	})
}

// GET FILM BY ID...
func (h *Handler) getFilmByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")

		return
	}

	film, err := h.services.FilmsList.GetByID(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, film)
}

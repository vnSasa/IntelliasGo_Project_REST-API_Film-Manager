package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	app "github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/model"
)

// CREATE DIRECTOR...

// @Summary Create director
// @Security ApiKeyAuth
// @Tags admin api directors
// @Description create director
// @ID create-director
// @Accept json
// @Produce json
// @Param input body app.DirectorsList true "director info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api-admin/directors/create [post]
func (h *Handler) createDiretor(c *gin.Context) {
	var input app.DirectorsList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	id, err := h.services.DirectorsList.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// GET ALL DIRECTORS...

type getAllDirectorsResponce struct {
	Directors []app.DirectorsList `json:"directors"`
}

// @Summary Get All Directors
// @Security ApiKeyAuth
// @Tags admin api directors
// @Description get all directors
// @ID get-all-directors
// @Accept json
// @Produce json
// @Success 200 {object} getAllDirectorsResponce
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api-admin/directors/all [get]
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

// @Summary Get Director By ID
// @Security ApiKeyAuth
// @Tags admin api directors
// @Description get director by id
// @ID get-director-by-id
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} app.DirectorsList
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api-admin/directors/{id} [get]
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

// UPDATE DIRECTOR...

// @Summary Update director
// @Security ApiKeyAuth
// @Tags admin api directors
// @Description update director
// @ID update-director
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param input body app.UpdateDirectorInput true "director info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api-admin/directors/{id} [put]
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

	if err := h.services.DirectorsList.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// DELETE DIRECTOR...

// @Summary Delete director
// @Security ApiKeyAuth
// @Tags admin api directors
// @Description delete director
// @ID delete-director
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api-admin/directors/{id} [delete]
func (h *Handler) deleteDiretor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")

		return
	}

	if err = h.services.DirectorsList.Delete(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// CREATE FILM...

// @Summary Create film
// @Security ApiKeyAuth
// @Tags admin api films
// @Description create film
// @ID create-film
// @Accept json
// @Produce json
// @Param input body app.FilmsList true "film info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api-admin/films/create [post]
func (h *Handler) createFilm(c *gin.Context) {
	var input app.FilmsList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	if err := h.validFilm(input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	id, err := h.services.FilmsList.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) validFilm(input app.FilmsList) error {
	_, err := h.services.DirectorsList.GetByID(input.DirectorID)
	if err != nil {
		return errors.New("director not found")
	}

	return nil
}

// GET ALL FILMS...

type getAllFilmsResponce struct {
	Films []app.FilmsList `json:"films"`
}

// @Summary Get All Films
// @Security ApiKeyAuth
// @Tags admin api films
// @Description get all films
// @ID get-all-films
// @Accept json
// @Produce json
// @Success 200 {object} getAllFilmsResponce
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api-admin/films/all [get]
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

// @Summary Get Film By ID
// @Security ApiKeyAuth
// @Tags admin api films
// @Description get film by id
// @ID get-film-by-id
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} app.FilmsList
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api-admin/films/{id} [get]
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

// UPDATE FILM...

// @Summary Update film
// @Security ApiKeyAuth
// @Tags admin api films
// @Description update film
// @ID update-film
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param input body app.UpdateFilmInput true "film info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api-admin/films/{id} [put]
func (h *Handler) updateFilm(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")

		return
	}

	var input app.UpdateFilmInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	if err := h.services.FilmsList.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// DELETE FILM...

// @Summary Delete film
// @Security ApiKeyAuth
// @Tags admin api films
// @Description delete film
// @ID delete-film
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param input body app.UpdateFilmInput true "film info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api-admin/films/{id} [delete]
func (h *Handler) deleteFilm(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")

		return
	}

	if err = h.services.FilmsList.Delete(id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

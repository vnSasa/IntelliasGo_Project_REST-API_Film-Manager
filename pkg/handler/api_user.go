package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
	"net/http"
	"strconv"
)

// GET ALL FILMS BY FILTERS...
func (h *Handler) getFilmsFilters(c *gin.Context) {
	var input app.FiltersFilmsInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	films, err := h.services.FilmsList.GetAllFilterFilms(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, getAllFilmsResponce{
		Films: films,
	})
}

// ADD FAVOURITE FILM...
func (h *Handler) addFavouriteFilm(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	filmID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")

		return
	}

	favouriteFilmID, err := h.services.FavouriteFilms.AddFavouriteFilm(userID, filmID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"favouriteFilmID": favouriteFilmID,
	})
}

// GET ALL FAVOURITE FILMS...
func (h *Handler) getAllFavouriteFilms(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	films, err := h.services.FavouriteFilms.GetAllFavouriteFilms(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, films)
}

// DELETE FILM FROM FAVOURITE...
func (h *Handler) deleteFavourite(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")

		return
	}

	err = h.services.FavouriteFilms.Delete(userID, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// ADD WISH FILM...
func (h *Handler) addWishFilm(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	filmID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid list id param")

		return
	}

	wishFilmID, err := h.services.WishFilms.AddWishFilm(userID, filmID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"wishFilmID": wishFilmID,
	})
}

// GET ALL WISH FILMS...
func (h *Handler) getAllWishFilms(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	films, err := h.services.WishFilms.GetAllWishFilms(userID)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, films)
}

// DELETE FILM FROM WISH...
func (h *Handler) deleteWish(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")

		return
	}

	err = h.services.WishFilms.Delete(userID, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

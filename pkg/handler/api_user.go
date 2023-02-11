package handler

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	app "github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/model"
)

// GET ALL FILMS BY FILTERS...

type getAllFilmsFilteredResponce struct {
	Films []app.FilmsList `json:"films"`
}

// @Summary Get Film By Filters
// @Security ApiKeyAuth
// @Tags user api films
// @Description get film by filters
// @ID get-film-by-filters
// @Accept json
// @Produce json
// @Param input body app.FiltersFilmsInput true "films info"
// @Success 200 {object} getAllFilmsFilteredResponce
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api-user/films/all [post].
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

	c.JSON(http.StatusOK, getAllFilmsFilteredResponce{
		Films: films,
	})
}

// EXPORT ALL FILMS TO CSV

// @Summary Export All Films To CSV
// @Security ApiKeyAuth
// @Tags user api films
// @Description export films
// @ID export-films
// @Accept json
// @Produce json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api-user/films/export [post].
func (h *Handler) exportFilmstoCSV(c *gin.Context) {
	films, err := h.services.FilmsList.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	csvfile, err := os.Create("films.csv")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	write := csv.NewWriter(csvfile)

	headers := []string{"list of films"}
	write.Write(headers)

	var id, name, genre, directorID, rate, year, minutes string

	for key := range films {
		list := make([]string, 0, 1+len(headers))

		id = fmt.Sprintf("| ID: %v", films[key].ID)
		name = fmt.Sprintf("| NAME: %v", films[key].Name)
		genre = fmt.Sprintf("| GENRE: %v", films[key].Genre)
		directorID = fmt.Sprintf("| DirectorID: %v", films[key].DirectorID)
		rate = fmt.Sprintf("| RATE: %v", films[key].Rate)
		year = fmt.Sprintf("| YEAR: %v", films[key].Year)
		minutes = fmt.Sprintf("| MINUTES: %v |", films[key].Minutes)

		list = append(
			list,
			id,
			name,
			genre,
			directorID,
			rate,
			year,
			minutes,
		)

		write.Write(list)
	}
	write.Flush()
	csvfile.Close()
}

// ADD FAVOURITE FILM...

// @Summary Add Favourite Film
// @Security ApiKeyAuth
// @Tags user api favourite
// @Description add favourite film
// @ID add-favourite-film
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api-user/films/favourite/{id}/add [post].
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

// @Summary Get All Favourite Film
// @Security ApiKeyAuth
// @Tags user api favourite
// @Description get all favourite film
// @ID get-all-favourite-film
// @Accept json
// @Produce json
// @Success 200 {integer} app.FilmsList
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api-user/films/favourite/all [get].
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

// @Summary Delete Favourite Film
// @Security ApiKeyAuth
// @Tags user api favourite
// @Description delete favourite film
// @ID delete-favourite-film
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api-user/films/favourite/{id} [delete].
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

// @Summary Add Wish Film
// @Security ApiKeyAuth
// @Tags user api wish
// @Description add wish film
// @ID add-wish-film
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api-user/films/wish/{id}/add [post].
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

// @Summary Get All Wish Film
// @Security ApiKeyAuth
// @Tags user api wish
// @Description get all wish film
// @ID get-all-wish-film
// @Accept json
// @Produce json
// @Success 200 {object} app.FilmsList
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api-user/films/wish/all [get].
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

// @Summary Delete Wish Film
// @Security ApiKeyAuth
// @Tags user api wish
// @Description delete wish film
// @ID delete-wish-film
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api-user/films/wish/{id} [delete].
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

// EXPORT FAVOURITE FILMS TO CSV

// @Summary Export Favourite Films To CSV
// @Security ApiKeyAuth
// @Tags user api favourite
// @Description export favourite films
// @ID export-favourite-films
// @Accept json
// @Produce json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api-user/films/favourite/export [post].
func (h *Handler) exportFtoCSV(c *gin.Context) {
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

	csvfile, err := os.Create("../favourite_films.csv")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	write := csv.NewWriter(csvfile)

	headers := []string{"list of my favourite films"}
	write.Write(headers)

	var id, name, genre, directorID, rate, year, minutes string

	for key := range films {
		list := make([]string, 0, 1+len(headers))

		id = fmt.Sprintf("| ID: %v", films[key].ID)
		name = fmt.Sprintf("| NAME: %v", films[key].Name)
		genre = fmt.Sprintf("| GENRE: %v", films[key].Genre)
		directorID = fmt.Sprintf("| DirectorID: %v", films[key].DirectorID)
		rate = fmt.Sprintf("| RATE: %v", films[key].Rate)
		year = fmt.Sprintf("| YEAR: %v", films[key].Year)
		minutes = fmt.Sprintf("| MINUTES: %v |", films[key].Minutes)

		list = append(
			list,
			id,
			name,
			genre,
			directorID,
			rate,
			year,
			minutes,
		)

		write.Write(list)
	}
	write.Flush()
	csvfile.Close()
}

// EXPORT WISH FILMS TO CSV

// @Summary Export Wish Films To CSV
// @Security ApiKeyAuth
// @Tags user api wish
// @Description export wish films
// @ID export-wish-films
// @Accept json
// @Produce json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api-user/films/wish/export [post].
func (h *Handler) exportWtoCSV(c *gin.Context) {
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

	csvfile, err := os.Create("wish_films.csv")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	write := csv.NewWriter(csvfile)

	headers := []string{"list of my wish films"}
	write.Write(headers)

	var id, name, genre, directorID, rate, year, minutes string

	for key := range films {
		list := make([]string, 0, 1+len(headers))

		id = fmt.Sprintf("| ID: %v", films[key].ID)
		name = fmt.Sprintf("| NAME: %v", films[key].Name)
		genre = fmt.Sprintf("| GENRE: %v", films[key].Genre)
		directorID = fmt.Sprintf("| DirectorID: %v", films[key].DirectorID)
		rate = fmt.Sprintf("| RATE: %v", films[key].Rate)
		year = fmt.Sprintf("| YEAR: %v", films[key].Year)
		minutes = fmt.Sprintf("| MINUTES: %v |", films[key].Minutes)

		list = append(
			list,
			id,
			name,
			genre,
			directorID,
			rate,
			year,
			minutes,
		)

		write.Write(list)
	}
	write.Flush()
	csvfile.Close()
}

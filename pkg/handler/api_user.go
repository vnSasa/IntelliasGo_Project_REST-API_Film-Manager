package handler

import (
	"net/http"
	"strconv"
	"os"
	"encoding/csv"
	"fmt"

	"github.com/gin-gonic/gin"
)

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

// EXPORT FAVOURITE FILMS TO CSV
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

	csvfile, err := os.Create("favourite_films.csv")
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
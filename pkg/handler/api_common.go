package handler

import (
	"net/http"
	"strconv"
	"os"
	"encoding/csv"
	"fmt"

	"github.com/gin-gonic/gin"
	app "github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/model"
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

// EXPORT ALL FILMS TO CSV
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

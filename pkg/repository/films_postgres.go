package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	app "github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/model"
)

type FilmsPostgres struct {
	db *sqlx.DB
}

func NewFilmPostgres(db *sqlx.DB) *FilmsPostgres {
	return &FilmsPostgres{db: db}
}

func (r *FilmsPostgres) Create(film app.FilmsList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var filmID int
	filmQuery := fmt.Sprintf("INSERT INTO %s (name, genre, director_id, rate, year, minutes) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", filmTable)
	row := tx.QueryRow(filmQuery, film.Name, film.Genre, film.DirectorID, fmt.Sprintf("%.1f", film.Rate), film.Year, film.Minutes)
	err = row.Scan(&filmID)
	if err != nil {
		if rb := tx.Rollback(); rb != nil {
			return 0, err
		}

		return 0, err
	}

	return filmID, tx.Commit()
}

func (r *FilmsPostgres) GetAll() ([]app.FilmsList, error) {
	var films []app.FilmsList
	query := fmt.Sprintf("SELECT id, name, genre, director_id, rate, year, minutes FROM %s", filmTable)
	err := r.db.Select(&films, query)

	return films, err
}

func (r *FilmsPostgres) GetAllFilterFilms(input app.FiltersFilmsInput) ([]app.FilmsList, error) {
	var films []app.FilmsList

	whereValue := make([]string, 0)
	orderByValue := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	var queryFilter, querySort, limit string

	if input.Genre != nil {
		whereValue = append(whereValue, fmt.Sprintf("genre=$%d", argID))
		args = append(args, *input.Genre)
		argID++
	}

	if input.MinRate != nil {
		whereValue = append(whereValue, fmt.Sprintf("rate >= $%d", argID))
		args = append(args, *input.MinRate)
		argID++
	}

	if input.SortRate != nil {
		orderByValue = append(orderByValue, "rate")
	}

	if input.SortYear != nil {
		orderByValue = append(orderByValue, "year")
	}

	if input.SortTime != nil {
		orderByValue = append(orderByValue, "minutes")
	}

	if input.Count != nil {
		limit = fmt.Sprintf("LIMIT $%d", argID)
		args = append(args, *input.Count)
	}

	whereQuery := strings.Join(whereValue, " and ")
	if len(whereQuery) != 0 {
		queryFilter = fmt.Sprintf("WHERE %s", whereQuery)
	}

	orderByQuery := strings.Join(orderByValue, ", ")
	if len(orderByQuery) != 0 {
		querySort = fmt.Sprintf("ORDER BY %s", orderByQuery)
	}

	query := fmt.Sprintf("SELECT id, name, genre, director_id, rate, year, minutes FROM %s %s %s %s", filmTable, queryFilter, querySort, limit)
	err := r.db.Select(&films, query, args...)

	return films, err
}

func (r *FilmsPostgres) GetByID(filmID int) (app.FilmsList, error) {
	var films app.FilmsList
	query := fmt.Sprintf("SELECT id, name, genre, director_id, rate, year, minutes FROM %s WHERE id = $1", filmTable)
	err := r.db.Get(&films, query, filmID)

	return films, err
}

func (r *FilmsPostgres) Update(filmID int, input app.UpdateFilmInput) error {
	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Name != nil {
		setValue = append(setValue, fmt.Sprintf("name=$%d", argID))
		args = append(args, *input.Name)
		argID++
	}

	if input.Genre != nil {
		setValue = append(setValue, fmt.Sprintf("genre=$%d", argID))
		args = append(args, *input.Genre)
		argID++
	}

	if input.DirectorID != nil {
		setValue = append(setValue, fmt.Sprintf("director_id=$%d", argID))
		args = append(args, *input.DirectorID)
		argID++
	}

	if input.Rate != nil {
		setValue = append(setValue, fmt.Sprintf("rate=$%d", argID))
		args = append(args, *input.Rate)
		argID++
	}

	if input.Year != nil {
		setValue = append(setValue, fmt.Sprintf("year=$%d", argID))
		args = append(args, *input.Year)
		argID++
	}

	if input.Minutes != nil {
		setValue = append(setValue, fmt.Sprintf("minutes=$%d", argID))
		args = append(args, *input.Minutes)
		argID++
	}

	setQuery := strings.Join(setValue, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", filmTable, setQuery, argID)

	args = append(args, filmID)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)

	return err
}

func (r *FilmsPostgres) Delete(filmID int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", filmTable)
	_, err := r.db.Exec(query, filmID)

	return err
}

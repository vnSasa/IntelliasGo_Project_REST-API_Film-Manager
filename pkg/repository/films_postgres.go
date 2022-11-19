package repository

import (
	"fmt"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
	"github.com/jmoiron/sqlx"
	"strings"
	"github.com/sirupsen/logrus"
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
	
	var filmId int
	filmQuery := fmt.Sprintf("INSERT INTO %s (name, genre, director_id, rate, year, minutes) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id", filmTable)
	row := tx.QueryRow(filmQuery, film.Name, film.Genre, film.DirectorId, film.Rate, film.Year, film.Minutes)
	err = row.Scan(&filmId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return filmId, tx.Commit()
}

func (r *FilmsPostgres) GetAll() ([]app.FilmsList, error) {
	var films []app.FilmsList
	query := fmt.Sprintf("SELECT id, name, genre, director_id, rate, year, minutes FROM %s", filmTable)
	err := r.db.Select(&films, query)

	return films, err
}

func (r *FilmsPostgres) GetById(filmId int) (app.FilmsList, error) {
	var films app.FilmsList
	query := fmt.Sprintf("SELECT id, name, genre, director_id, rate, year, minutes FROM %s WHERE id = $1", filmTable)
	err := r.db.Get(&films, query, filmId)

	return films, err
}

func (r *FilmsPostgres) Update(filmId int, input app.UpdateFilmInput) error {
	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValue = append(setValue, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Genre != nil {
		setValue = append(setValue, fmt.Sprintf("genre=$%d", argId))
		args = append(args, *input.Genre)
		argId++
	}

	if input.DirectorId != nil {
		setValue = append(setValue, fmt.Sprintf("director_id=$%d", argId))
		args = append(args, *input.DirectorId)
		argId++
	}

	if input.Rate != nil {
		setValue = append(setValue, fmt.Sprintf("rate=$%d", argId))
		args = append(args, *input.Rate)
		argId++
	}

	if input.Year != nil {
		setValue = append(setValue, fmt.Sprintf("year=$%d", argId))
		args = append(args, *input.Year)
		argId++
	}

	if input.Minutes != nil {
		setValue = append(setValue, fmt.Sprintf("minutes=$%d", argId))
		args = append(args, *input.Minutes)
		argId++
	}

	setQuery := strings.Join(setValue, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", filmTable, setQuery, argId)

	args = append(args, filmId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *FilmsPostgres) Delete(filmId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", filmTable)
	_, err := r.db.Exec(query, filmId)

	return err
}
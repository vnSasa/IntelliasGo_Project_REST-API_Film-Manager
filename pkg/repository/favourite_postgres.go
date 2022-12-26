package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
)

type FavouriteFilmsPostgres struct {
	db *sqlx.DB
}

func NewFavouriteFilmsPostgres(db *sqlx.DB) *FavouriteFilmsPostgres {
	return &FavouriteFilmsPostgres{db: db}
}

func (r *FavouriteFilmsPostgres) AddFavouriteFilm(userID, filmID int) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var favouriteFilmID int
	addFavouriteFilm := fmt.Sprintf("INSERT INTO %s (user_id, film_id) VALUES ($1, $2) RETURNING id", favouriteTable)

	row := tx.QueryRow(addFavouriteFilm, userID, filmID)
	err = row.Scan(&favouriteFilmID)
	if err != nil {
		if rb := tx.Rollback(); rb != nil {
			return 0, err
		}

		return 0, err
	}

	return favouriteFilmID, tx.Commit()
}

func (r *FavouriteFilmsPostgres) GetAllFavouriteFilms(userID int) ([]app.FilmsList, error) {
	var films []app.FilmsList
	query := fmt.Sprintf(`SELECT f.id, f.name, f.genre, f.director_id, f.rate, f.year, f.minutes FROM %s f
		INNER JOIN %s ff on f.id = ff.film_id WHERE ff.user_id = $1`, filmTable, favouriteTable)
	err := r.db.Select(&films, query, userID)

	return films, err
}

func (r *FavouriteFilmsPostgres) Delete(userID, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND film_id = $2", favouriteTable)
	_, err := r.db.Exec(query, userID, id)

	return err
}

package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	app "github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
)

type WishFilmsPostgres struct {
	db *sqlx.DB
}

func NewWishFilmsPostgres(db *sqlx.DB) *WishFilmsPostgres {
	return &WishFilmsPostgres{db: db}
}

func (r *WishFilmsPostgres) AddWishFilm(userID, filmID int) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var wishFilmID int
	addWishFilm := fmt.Sprintf("INSERT INTO %s (user_id, film_id) VALUES ($1, $2) RETURNING id", wishTable)

	row := tx.QueryRow(addWishFilm, userID, filmID)
	err = row.Scan(&wishFilmID)
	if err != nil {
		if rb := tx.Rollback(); rb != nil {
			return 0, err
		}

		return 0, err
	}

	return wishFilmID, tx.Commit()
}

func (r *WishFilmsPostgres) GetAllWishFilms(userID int) ([]app.FilmsList, error) {
	var films []app.FilmsList
	query := fmt.Sprintf(`SELECT f.id, f.name, f.genre, f.director_id, f.rate, f.year, f.minutes FROM %s f
		INNER JOIN %s wf on f.id = wf.film_id WHERE wf.user_id = $1`, filmTable, wishTable)
	err := r.db.Select(&films, query, userID)

	return films, err
}

func (r *WishFilmsPostgres) Delete(userID, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND film_id = $2", wishTable)
	_, err := r.db.Exec(query, userID, id)

	return err
}

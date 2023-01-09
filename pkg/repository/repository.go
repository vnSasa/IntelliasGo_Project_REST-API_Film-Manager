package repository

import (
	"github.com/jmoiron/sqlx"
	app "github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
)

type Authorization interface {
	CreateAdmin(admin app.User) (int, error)
	CreateUser(user app.User) (int, error)
	DeleteUser(id int) error
	GetUser(login, password string) (app.User, error)
}

type DirectorsList interface {
	Create(director app.DirectorsList) (int, error)
	GetAll() ([]app.DirectorsList, error)
	GetByID(directorID int) (app.DirectorsList, error)
	Update(directorID int, input app.UpdateDirectorInput) error
	Delete(directorID int) error
}

type FilmsList interface {
	Create(film app.FilmsList) (int, error)
	GetAll() ([]app.FilmsList, error)
	GetAllFilterFilms(input app.FiltersFilmsInput) ([]app.FilmsList, error)
	GetByID(filmID int) (app.FilmsList, error)
	Update(filmID int, input app.UpdateFilmInput) error
	Delete(filmID int) error
}

type FavouriteFilms interface {
	AddFavouriteFilm(userID, filmID int) (int, error)
	GetAllFavouriteFilms(userID int) ([]app.FilmsList, error)
	Delete(userID, id int) error
}

type WishFilms interface {
	AddWishFilm(userID, filmID int) (int, error)
	GetAllWishFilms(userID int) ([]app.FilmsList, error)
	Delete(userID, id int) error
}

type Repository struct {
	Authorization
	DirectorsList
	FilmsList
	FavouriteFilms
	WishFilms
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization:  NewAuthPostgres(db),
		DirectorsList:  NewDirectorPostgres(db),
		FilmsList:      NewFilmPostgres(db),
		FavouriteFilms: NewFavouriteFilmsPostgres(db),
		WishFilms:      NewWishFilmsPostgres(db),
	}
}

package repository

import (
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateAdmin(admin app.User) (int, error)
	CreateUser(user app.User) (int, error)
	DeleteUser(user app.User) (int, error)
	GetUser(login, password string) (app.User, error)
	GetUserById(id int) (app.User, error)
}

type DirectorsList interface {
	Create(director app.DirectorsList) (int, error)
	GetAll() ([]app.DirectorsList, error)
	GetById(directorId int) (app.DirectorsList, error)
	Update(directorId int, input app.UpdateDirectorInput) error
	Delete(directorId int) error
}

type FilmsList interface {
	Create(film app.FilmsList) (int, error)
	GetAll() ([]app.FilmsList, error)
	GetById(filmId int) (app.FilmsList, error)
	Update(filmId int, input app.UpdateFilmInput) error
	Delete(filmId int) error
}

type FavouriteFilms interface {

}

type WishFilms interface {

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
		Authorization: NewAuthPostgres(db),
		DirectorsList: NewDirectorPostgres(db),
		FilmsList: NewFilmPostgres(db),
	}
}
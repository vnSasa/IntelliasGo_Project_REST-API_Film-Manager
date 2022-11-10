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

type FilmsList interface {

}

type DirectorsFilms interface {

}

type FavouriteFilms interface {

}

type WishFilms interface {

}

type Repository struct {
	Authorization
	FilmsList
	DirectorsFilms
	FavouriteFilms
	WishFilms
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
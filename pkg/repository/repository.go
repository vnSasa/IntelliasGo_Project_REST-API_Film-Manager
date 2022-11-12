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

type DirectorList interface {
	Create(userLogin string, director app.DirectorList) (int, error)
}

type FilmsList interface {

}

type FavouriteFilms interface {

}

type WishFilms interface {

}

type Repository struct {
	Authorization
	DirectorList
	FilmsList
	FavouriteFilms
	WishFilms
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		DirectorList: NewDirectorPostgres(db),
	}
}
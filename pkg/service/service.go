package service

import (
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/pkg/repository"
)

type Authorization interface {
	CreateAdmin(admin app.User) (int, error)
	CreateUser(user app.User) (int, error)
	GetUserById(id int) (app.User, error)
	DeleteUser(user app.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, string, error)
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

type Service struct {
	Authorization
	DirectorsList
	FilmsList
	FavouriteFilms
	WishFilms
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		DirectorsList: NewDirectorService(repos.DirectorsList),
		FilmsList: NewFilmsService(repos.FilmsList),
	}
}
package service

import (
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/pkg/repository"
)

type Authorization interface {
	CreateAdmin(admin app.User) (int, error)
	CreateUser(user app.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(token string) (int, string, error)
}

type FilmsList interface {

}

type DirectorsFilms interface {

}

type FavouriteFilms interface {

}

type WishFilms interface {

}

type Service struct {
	Authorization
	FilmsList
	DirectorsFilms
	FavouriteFilms
	WishFilms
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
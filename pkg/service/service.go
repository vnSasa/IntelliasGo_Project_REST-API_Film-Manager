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

type DirectorList interface {
	Create(userLogin string, director app.DirectorList) (int, error)
}

type FilmsList interface {

}

type FavouriteFilms interface {

}

type WishFilms interface {

}

type Service struct {
	Authorization
	DirectorList
	FilmsList
	FavouriteFilms
	WishFilms
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		DirectorList: NewDirectorService(repos.DirectorList),
	}
}
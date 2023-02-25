package service

import (
	app "github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/model"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/pkg/repository"
)

//go:generate go run github.com/golang/mock/mockgen -package mocks -destination=mocks/mock_service.go -source=service.go

type Authorization interface {
	CreateAdmin(admin app.User) (int, error)
	CreateUser(user app.User) (int, error)
	GetUser(login, password string) error
	DeleteUser(id int) error
	GenerateToken(login, password string) (*app.TokenDetails, error)
	RefreshToken(refreshData *app.RefreshTokenClaims) (*app.TokenDetails, error)
	ParseToken(token string) (*app.AccessTokenClaims, error)
	VerifyAdminToken(accessToken string) (*app.AccessTokenClaims, error)
	VerifyUserToken(accessToken string) (*app.AccessTokenClaims, error)
	ParseRefreshToken(refreshToken string) (*app.RefreshTokenClaims, error)
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

type Service struct {
	Authorization
	DirectorsList
	FilmsList
	FavouriteFilms
	WishFilms
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization:  NewAuthService(repos.Authorization),
		DirectorsList:  NewDirectorService(repos.DirectorsList),
		FilmsList:      NewFilmsService(repos.FilmsList),
		FavouriteFilms: NewFavouriteFilmsService(repos.FavouriteFilms),
		WishFilms:      NewWishFilmsService(repos.WishFilms),
	}
}

package service

import (
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/pkg/repository"
)

type FavouriteFilmsService struct {
	repo repository.FavouriteFilms
}

func NewFavouriteFilmsService(repo repository.FavouriteFilms) *FavouriteFilmsService {
	return &FavouriteFilmsService{repo: repo}
}

func (s *FavouriteFilmsService) AddFavouriteFilm(userID, filmID int) (int, error) {
	return s.repo.AddFavouriteFilm(userID, filmID)
}

func (s *FavouriteFilmsService) GetAllFavouriteFilms(userID int) ([]app.FilmsList, error) {
	return s.repo.GetAllFavouriteFilms(userID)
}

func (s *FavouriteFilmsService) Delete(userID, id int) error {
	return s.repo.Delete(userID, id)
}

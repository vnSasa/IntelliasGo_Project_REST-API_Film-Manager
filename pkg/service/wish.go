package service

import (
	app "github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/pkg/repository"
)

type WishFilmsService struct {
	repo repository.WishFilms
}

func NewWishFilmsService(repo repository.WishFilms) *WishFilmsService {
	return &WishFilmsService{repo: repo}
}

func (s *WishFilmsService) AddWishFilm(userID, filmID int) (int, error) {
	return s.repo.AddWishFilm(userID, filmID)
}

func (s *WishFilmsService) GetAllWishFilms(userID int) ([]app.FilmsList, error) {
	return s.repo.GetAllWishFilms(userID)
}

func (s *WishFilmsService) Delete(userID, id int) error {
	return s.repo.Delete(userID, id)
}

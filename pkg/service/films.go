package service

import (
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/pkg/repository"
)

type FilmsService struct {
	repo repository.FilmsList
}

func NewFilmsService(repo repository.FilmsList) *FilmsService {
	return &FilmsService{repo: repo}
}

func (s *FilmsService) Create(film app.FilmsList) (int, error) {
	return s.repo.Create(film)
}

func (s *FilmsService) GetAll() ([]app.FilmsList, error) {
	return s.repo.GetAll()
}

func (s *FilmsService) GetById(filmId int) (app.FilmsList, error) {
	return s.repo.GetById(filmId)
}

func (s *FilmsService) Update(filmId int, input app.UpdateFilmInput) error {
	return s.repo.Update(filmId, input)
}

func (s *FilmsService) Delete(filmId int) error {
	return s.repo.Delete(filmId)
}
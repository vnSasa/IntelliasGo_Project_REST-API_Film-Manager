package service

import (
	app "github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
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

func (s *FilmsService) GetAllFilterFilms(input app.FiltersFilmsInput) ([]app.FilmsList, error) {
	return s.repo.GetAllFilterFilms(input)
}

func (s *FilmsService) GetByID(filmID int) (app.FilmsList, error) {
	return s.repo.GetByID(filmID)
}

func (s *FilmsService) Update(filmID int, input app.UpdateFilmInput) error {
	return s.repo.Update(filmID, input)
}

func (s *FilmsService) Delete(filmID int) error {
	return s.repo.Delete(filmID)
}

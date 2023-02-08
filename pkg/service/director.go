package service

import (
	app "github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/model"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/pkg/repository"
)

type DirectorService struct {
	repo repository.DirectorsList
}

func NewDirectorService(repo repository.DirectorsList) *DirectorService {
	return &DirectorService{repo: repo}
}

func (s *DirectorService) Create(director app.DirectorsList) (int, error) {
	return s.repo.Create(director)
}

func (s *DirectorService) GetAll() ([]app.DirectorsList, error) {
	return s.repo.GetAll()
}

func (s *DirectorService) GetByID(directorID int) (app.DirectorsList, error) {
	return s.repo.GetByID(directorID)
}

func (s *DirectorService) Update(directorID int, input app.UpdateDirectorInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(directorID, input)
}

func (s *DirectorService) Delete(directorID int) error {
	return s.repo.Delete(directorID)
}

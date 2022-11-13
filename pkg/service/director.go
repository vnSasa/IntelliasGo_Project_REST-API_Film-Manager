package service

import (
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager"
	"github.com/vnSasa/IntelliasGo_Project_REST-API_Film-Manager/pkg/repository"
)

type DirectorService struct {
	repo repository.DirectorList
}

func NewDirectorService(repo repository.DirectorList) *DirectorService {
	return &DirectorService{repo: repo}
}

func (s *DirectorService) Create(director app.DirectorList) (int, error) {
	return s.repo.Create(director)
}

func (s *DirectorService) GetAll() ([]app.DirectorList, error) {
	return s.repo.GetAll()
}

func (s *DirectorService) GetById(directorId int) (app.DirectorList, error) {
	return s.repo.GetById(directorId)
}

func (s *DirectorService) Update(directorId int, input app.UpdateDirectorInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(directorId, input)
}

func (s *DirectorService) Delete(directorId int) error {
	return s.repo.Delete(directorId)
}
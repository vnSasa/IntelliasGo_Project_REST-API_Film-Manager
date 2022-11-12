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

func (s *DirectorService) Create(userLogin string, director app.DirectorList) (int, error) {
	return s.repo.Create(userLogin, director)
}
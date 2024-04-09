package service

import (
	"fmt"

	"github.com/sixojke/lolz-test/internal/domain"
	"github.com/sixojke/lolz-test/internal/repository"
)

type GenreService struct {
	repo repository.Genre
}

func NewGenreService(repo repository.Genre) *GenreService {
	return &GenreService{repo: repo}
}

func (s *GenreService) GetAll() (*[]domain.Genre, error) {
	genres, err := s.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("genre service get all: %v", err)
	}

	return genres, nil
}

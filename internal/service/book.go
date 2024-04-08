package service

import (
	"fmt"

	"github.com/sixojke/lolz-test/internal/domain"
	"github.com/sixojke/lolz-test/internal/repository"
)

type BookService struct {
	repo repository.Book
}

func NewBookService() *BookService {
	return &BookService{}
}

func (s *BookService) Create(inp *domain.BookCreateInp) error {
	if err := s.repo.Create(inp); err != nil {
		return fmt.Errorf("book service create: %v", err)
	}

	return nil
}

func (s *BookService) GetById(inp *domain.BookGetByIdInp) (*domain.BookGetByIdOut, error) {
	out, err := s.repo.GetById(inp)
	if err != nil {
		return nil, fmt.Errorf("book service get by id: %v", err)
	}

	return out, nil
}

func (s *BookService) Delete(inp *domain.BookDeleteInp) error {
	if err := s.repo.Delete(inp); err != nil {
		return fmt.Errorf("book service delete: %v", err)
	}

	return nil
}

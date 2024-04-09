package service

import (
	"fmt"

	"github.com/sixojke/lolz-test/internal/domain"
	"github.com/sixojke/lolz-test/internal/repository"
)

type BookService struct {
	repo repository.Book
}

func NewBookService(repo repository.Book) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) Create(inp *domain.BookCreateInp) error {
	if err := s.repo.Create(inp); err != nil {
		return fmt.Errorf("book service create: %v", err)
	}

	return nil
}

func (s *BookService) GetById(inp *domain.BookGetByIdInp) (*domain.BookGetByIdOut, error) {
	book, err := s.repo.GetById(inp)
	if err != nil {
		return nil, fmt.Errorf("book service get by id: %v", err)
	}

	return book, nil
}

func (s *BookService) GetByGenre(inp *domain.BooksGetByGenreInp) (*domain.PaginationOut, error) {
	books, err := s.repo.GetByGenre(inp)
	if err != nil {
		return nil, fmt.Errorf("book service get by genre: %v", err)
	}

	return &domain.PaginationOut{
		Limit:  inp.Limit,
		Offset: inp.Offset,
		Data:   books,
	}, nil
}

func (s *BookService) Delete(inp *domain.BookDeleteInp) error {
	if err := s.repo.Delete(inp); err != nil {
		return fmt.Errorf("book service delete: %v", err)
	}

	return nil
}

func (s *BookService) Search(inp *domain.BooksSearchInp) (*domain.PaginationOut, error) {
	books, err := s.repo.Search(inp)
	if err != nil {
		return nil, fmt.Errorf("book service search: %v", err)
	}

	return &domain.PaginationOut{
		Data:   books,
		Limit:  inp.Limit,
		Offset: inp.Offset,
	}, nil
}

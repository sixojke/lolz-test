package service

import (
	"github.com/sixojke/lolz-test/internal/domain"
	"github.com/sixojke/lolz-test/internal/repository"
)

type Book interface {
	Create(inp *domain.BookCreateInp) error
	GetById(inp *domain.BookGetByIdInp) (*domain.BookGetByIdOut, error)
	GetByGenre(inp *domain.BooksGetByGenreInp) (*domain.PaginationOut, error)
	Delete(inp *domain.BookDeleteInp) error
	Search(inp *domain.BooksSearchInp) (*domain.PaginationOut, error)
}

type Genre interface {
	GetAll() (*[]domain.Genre, error)
}

type Deps struct {
	Repo *repository.Repository
}

type Service struct {
	Book  Book
	Genre Genre
}

func NewService(deps *Deps) *Service {
	return &Service{
		Book:  NewBookService(deps.Repo.Book),
		Genre: NewGenreService(deps.Repo.Genre),
	}
}

package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sixojke/lolz-test/internal/domain"
)

type Book interface {
	Create(inp *domain.BookCreateInp) error
	GetById(inp *domain.BookGetByIdInp) (*domain.BookGetByIdOut, error)
	GetByGenre(inp *domain.BooksGetByGenreInp) (*[]domain.BooksGetByGenreOut, error)
	Delete(inp *domain.BookDeleteInp) error
	Search(inp *domain.BooksSearchInp) (*[]domain.BooksSearchOut, error)
}

type Genre interface {
	GetAll() (*[]domain.Genre, error)
}

type Deps struct {
	Postgres *sqlx.DB
}

type Repository struct {
	Book  Book
	Genre Genre
}

func NewRepository(deps *Deps) *Repository {
	return &Repository{
		Book:  NewBookPostgres(deps.Postgres),
		Genre: NewGenrePostgres(deps.Postgres),
	}
}

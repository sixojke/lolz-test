package service

import "github.com/sixojke/lolz-test/internal/domain"

type Book interface {
	Create(inp *domain.BookCreateInp) error
	GetById(inp *domain.BookGetByIdInp) (*domain.BookGetByIdOut, error)
	Delete(inp *domain.BookDeleteInp) error
}

type Deps struct {
}

type Service struct {
	Book Book
}

func NewService() *Service {
	return &Service{}
}

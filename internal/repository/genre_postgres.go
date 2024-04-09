package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sixojke/lolz-test/internal/domain"
)

type GenrePostgres struct {
	db *sqlx.DB
}

func NewGenrePostgres(db *sqlx.DB) *GenrePostgres {
	return &GenrePostgres{db: db}
}

func (r *GenrePostgres) GetAll() (*[]domain.Genre, error) {
	var genres []domain.Genre
	query := fmt.Sprintf(`SELECT id, name FROM %s`, genre)
	if err := r.db.Select(&genres, query); err != nil {
		return nil, fmt.Errorf("error get all genres: %v", err)
	}

	return &genres, nil
}

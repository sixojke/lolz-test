package repository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sixojke/lolz-test/internal/domain"
)

type BookPostgres struct {
	db *sqlx.DB
}

func NewBookPostgres(db *sqlx.DB) *BookPostgres {
	return &BookPostgres{db: db}
}

func (r *BookPostgres) Create(inp *domain.BookCreateInp) error {
	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error create tx: %v", err)
	}

	// сканируем id жанра, который будет присвоин книге
	var genre_id string
	query := fmt.Sprintf(`SELECT id FROM %s WHERE name = $1`, genre)
	if err := tx.QueryRow(query, inp.Genre).Scan(&genre_id); err != nil {
		// добавляем жанр пользователя, если его нет в базе
		if err != sql.ErrNoRows {
			tx.Rollback()
			return fmt.Errorf("error select genre_id: %v", err)
		}

		query = fmt.Sprintf(`INSERT INTO %s (name) VALUES ($1) RETURNING id`, genre)
		if err := tx.QueryRow(query, inp.Genre).Scan(&genre_id); err != nil {
			tx.Rollback()
			return fmt.Errorf("error insert new genre: %v", err)
		}
	}

	// создание книги
	query = fmt.Sprintf(`INSERT INTO %s (name, author, description, genre_id) VALUES ($1, $2, $3)`, book)
	if _, err := tx.Exec(query, inp.Name, inp.Author, inp.Description, genre_id); err != nil {
		tx.Rollback()
		return fmt.Errorf("error insert book: %v", err)
	}

	return nil
}

func (r *BookPostgres) GetById(inp *domain.BookGetByIdInp) (*domain.BookGetByIdOut, error) {
	query := fmt.Sprintf(`
		SELECT book.name, book.description, book.genre_id, genre.name 
		FROM %s 
		INNER JOIN %s ON book.genre_id = genre.id
		WHERE book.id = $1`, book, genre)

	var out domain.BookGetByIdOut
	if err := r.db.QueryRow(query, inp.Id).Scan(&out); err != nil {
		return nil, fmt.Errorf("error get book by id: %v", err)
	}

	return &out, nil
}

func (r *BookPostgres) Delete(inp *domain.BookDeleteInp) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, book)
	if _, err := r.db.Exec(query, inp.Id); err != nil {
		return fmt.Errorf("error delete book: %v", err)
	}

	return nil
}

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
	query = fmt.Sprintf(`INSERT INTO %s (name, author, description, genre_id) VALUES ($1, $2, $3, $4)`, book)
	if _, err := tx.Exec(query, inp.Name, inp.Author, inp.Description, genre_id); err != nil {
		tx.Rollback()
		return fmt.Errorf("error insert book: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error tx.Commit: %v", err)
	}

	return nil
}

func (r *BookPostgres) GetById(inp *domain.BookGetByIdInp) (*domain.BookGetByIdOut, error) {
	query := fmt.Sprintf(`
		SELECT book.name, book.description, genre.name 
		FROM %s 
		INNER JOIN %s ON book.genre_id = genre.id
		WHERE book.id = $1`, book, genre)

	var book domain.BookGetByIdOut
	if err := r.db.QueryRow(query, inp.Id).Scan(&book.Name, &book.Description, &book.Genre); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error get book by id: %v", err)
	}

	return &book, nil
}

func (r *BookPostgres) GetByGenre(inp *domain.BooksGetByGenreInp) (*[]domain.BooksGetByGenreOut, error) {
	var books []domain.BooksGetByGenreOut

	var query string
	if inp.Genre == "all" {
		query = fmt.Sprintf(`SELECT id, author, name
			FROM %s 
			ORDER BY book.id, book.name, book.author
			OFFSET $1
			LIMIT $2`, book)

		if err := r.db.Select(&books, query, inp.Offset, inp.Limit); err != nil {
			return nil, fmt.Errorf("error get all books by genre: %v", err)
		}
	} else {
		query = fmt.Sprintf(`SELECT book.id, book.author, book.name FROM %s 
			JOIN %s ON book.genre_id = genre.id 
			WHERE genre.name = $1 
			ORDER BY book.name
			OFFSET $2
			LIMIT $3`, book, genre)

		if err := r.db.Select(&books, query, inp.Genre, inp.Offset, inp.Limit); err != nil {
			return nil, fmt.Errorf("error get books by genre: %v", err)
		}
	}

	return &books, nil
}

func (r *BookPostgres) Delete(inp *domain.BookDeleteInp) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, book)
	if _, err := r.db.Exec(query, inp.Id); err != nil {
		return fmt.Errorf("error delete book: %v", err)
	}

	return nil
}

func (r *BookPostgres) Search(inp *domain.BooksSearchInp) (*[]domain.BooksSearchOut, error) {
	fmt.Println(inp)
	var books []domain.BooksSearchOut
	query := fmt.Sprintf(`SELECT id, name, author FROM %s `, book) +
		`WHERE author LIKE '%' || $1 || '%' OR name LIKE '%' || $1 || '%'
        ORDER BY name
        OFFSET $2
        LIMIT $3;`

	if err := r.db.Select(&books, query, inp.String, inp.Offset, inp.Limit); err != nil {
		return nil, fmt.Errorf("error select books by string: %v", err)
	}

	return &books, nil
}

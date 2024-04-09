package domain

import (
	"fmt"
	"strings"
)

type BookCreateInp struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
}

func (d *BookCreateInp) Validate() error {
	var invalidData []string
	if d.Name == "" {
		invalidData = append(invalidData, "name")
	}

	if d.Author == "" {
		invalidData = append(invalidData, "author")
	}

	if d.Genre == "" {
		invalidData = append(invalidData, "genre")
	}

	if len(invalidData) != 0 {
		return fmt.Errorf("invalid data: " + strings.Join(invalidData, ", "))
	}

	return nil
}

type BookGetByIdInp struct {
	Id string `json:"id"`
}

func (d *BookGetByIdInp) Validate() error {
	if d.Id == "" {
		return fmt.Errorf("invalid data: id")
	}

	return nil
}

type BookGetByIdOut struct {
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Genre       string `json:"genre" db:"genre"`
}

type BooksGetByGenreInp struct {
	Genre  string `json:"genre"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

func (d *BooksGetByGenreInp) Validate(defaultLimit, defaultOffset int) {
	if d.Genre == "" {
		d.Genre = "all"
	}

	if d.Limit == 0 {
		d.Limit = defaultLimit
	}

	if d.Offset == 0 {
		d.Offset = defaultOffset
	}
}

type BooksGetByGenreOut struct {
	Id     string `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Author string `json:"author" db:"author"`
}

type BookDeleteInp struct {
	Id string `json:"id"`
}

func (d *BookDeleteInp) Validate() error {
	if d.Id == "" {
		return fmt.Errorf("invalid data: id")
	}

	return nil
}

type BooksSearchInp struct {
	String string `json:"string"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

func (d *BooksSearchInp) Validate(defaultLimit, defaultOffset int) {
	if d.Limit == 0 {
		d.Limit = defaultLimit
	}

	if d.Offset == 0 {
		d.Offset = defaultOffset
	}
}

type BooksSearchOut struct {
	Id     string `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Author string `json:"author" db:"author"`
}

type Genre struct {
	Id   string `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type PaginationOut struct {
	Limit  int         `json:"limit"`
	Offset int         `json:"offset"`
	Data   interface{} `json:"data"`
}

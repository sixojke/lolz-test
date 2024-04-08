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
	Name        string `json:"name"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
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

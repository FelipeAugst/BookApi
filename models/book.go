package models

import (
	"errors"
	"strings"
)

type Book struct {
	ID     uint64 `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   uint   `json:"year"`
	Status string `json:"status"`
}

func (b *Book) Validate() error {

	switch {
	case b.Title == "":
		return errors.New("empty title")
	case b.Author == "":
		return errors.New("empty author")
	default:
		b.Title = strings.TrimSpace(b.Title)
		b.Author = strings.TrimSpace(b.Author)
		return nil
	}
}

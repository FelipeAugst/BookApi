package db

import (
	"bookapi/models"
	"errors"
	"slices"
)

type Db struct {
	currentID uint64
	data      []models.Book
}

func (d *Db) Insert(book models.Book) {
	book.ID = d.currentID
	d.data = append(d.data, book)
	d.currentID++
}

func (d *Db) List(pagination bool, idx, limit int) ([]models.Book, error) {
	if !pagination {
		return d.data, nil
	}
	if idx < 0 || limit > len(d.data) {
		return nil, errors.New("Invalid page interval")
	}

	return d.data[idx:limit], nil
}

func (d *Db) SearchByAuthor(search string, pagination bool, idx, limit int) ([]models.Book, error) {

	if idx < 0 || limit > len(d.data) {
		return nil, errors.New("Invalid page interval")
	}

	var results []models.Book

	for _, book := range d.data {
		if book.Author == search {
			results = append(results, book)
		}

	}

	if !pagination {
		return results, nil
	}

	return results[idx:limit], nil
}

func (d *Db) SearchByTitle(search string, pagination bool, idx, limit int) ([]models.Book, error) {

	if idx < 0 || limit > len(d.data) {
		return nil, errors.New("Invalid page interval")
	}

	var results []models.Book

	for _, book := range d.data {
		if book.Title == search {
			results = append(results, book)
		}

	}

	if !pagination {
		return results, nil
	}

	return results[idx:limit], nil
}

func (d *Db) Update(id uint64, book models.Book) {
	for idx, _ := range d.data {
		if d.data[idx].ID == id {
			d.data[id] = book
			return
		}
	}

}

func (d *Db) Delete(id uint64, book models.Book) {
	for idx, _ := range d.data {
		if d.data[idx].ID == id {
			d.data = slices.Delete[[]models.Book](d.data, idx, idx)
			return
		}
	}

}

var DefaultDb = Db{currentID: 0, data: make([]models.Book, 10)}

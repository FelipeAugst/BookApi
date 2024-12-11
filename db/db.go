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

func (d *Db) Insert(book models.Book) uint64 {
	book.ID = d.currentID
	d.data = append(d.data, book)
	d.currentID++
	return book.ID
}

func (d *Db) List(pagination bool, page, limit uint64) ([]models.Book, error) {
	if !pagination {
		return d.data, nil
	}
	if page*limit > d.currentID {
		return nil, errors.New("invalid page interval")
	}
	return d.data, nil
}

func (d *Db) SearchByAuthor(search string, pagination bool, page, limit uint64) ([]models.Book, error) {

	if page*limit > d.currentID {
		return nil, errors.New("invalid page interval")
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

	return d.data[page*limit : limit], nil
}

func (d *Db) SearchByTitle(search string, pagination bool, idx, limit uint64) ([]models.Book, error) {

	if limit > d.currentID {
		return nil, errors.New("invalid page interval")
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

func (d *Db) Delete(id uint64) {
	for idx, _ := range d.data {
		if d.data[idx].ID == id {
			d.data = slices.DeleteFunc[[]models.Book](d.data, func(b models.Book) bool { return b.ID == id })
			return
		}
	}

}

func NewDb() *Db {
	d := new(Db)
	d.currentID = 0
	d.data = make([]models.Book, 0)
	return d
}

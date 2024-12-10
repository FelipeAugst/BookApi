package service

import (
	"bookapi/models"
	"bookapi/service/db"
)

type BookService interface {
	Create(models.Book) error
	List(bool, int, int) ([]models.Book, error)
	Update(models.Book)
	ByAuthor(string, bool, int, int) ([]models.Book, error)
	ByTitle(string, bool, int, int) ([]models.Book, error)
	Delete(id uint)
}

type bookServiceImpl struct {
	dataBase  db.Db
	currentId uint64
}

func (b *bookServiceImpl) Create(book models.Book) error {
	return b.Create(book)
}

func (b *bookServiceImpl) Delete(id uint) {
	b.Delete(id)
}

func (b *bookServiceImpl) List(pagination bool, idx, limit int) ([]models.Book, error) {
	search, err := b.dataBase.List(pagination, idx, limit)
	if err != nil {
		return nil, err
	}
	return search, nil
}

func (b *bookServiceImpl) ByAuthor(author string, pagination bool, idx, limit int) ([]models.Book, error) {
	return b.dataBase.SearchByAuthor(author, pagination, idx, limit)
}

func (b *bookServiceImpl) ByTitle(title string, pagination bool, idx, limit int) ([]models.Book, error) {
	return b.dataBase.SearchByAuthor(title, pagination, idx, limit)
}

func (b *bookServiceImpl) Update(book models.Book) {
	b.dataBase.Update(book.ID, book)
}

func NewBookService() BookService {
	bs := new(bookServiceImpl)
	bs.dataBase = db.DefaultDb

	return bs
}

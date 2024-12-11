package service

import (
	"bookapi/db"
	"bookapi/models"
)

var (
	DefaultDb = db.NewDb()
)

type BookService interface {
	Create(models.Book) (uint64, error)
	List(bool, uint64, uint64) ([]models.Book, error)
	Update(models.Book)
	ByAuthor(string, bool, uint64, uint64) ([]models.Book, error)
	ByTitle(string, bool, uint64, uint64) ([]models.Book, error)
	Delete(id uint64)
}

type bookServiceImpl struct {
	dataBase *db.Db
}

func (b *bookServiceImpl) Create(book models.Book) (uint64, error) {
	id := b.dataBase.Insert(book)
	return id, nil
}

func (b *bookServiceImpl) Delete(id uint64) {
	b.dataBase.Delete(id)
}

func (b *bookServiceImpl) List(pagination bool, idx, limit uint64) ([]models.Book, error) {
	search, err := b.dataBase.List(pagination, idx, limit)
	if err != nil {
		return nil, err
	}
	return search, nil
}

func (b *bookServiceImpl) ByAuthor(author string, pagination bool, idx, limit uint64) ([]models.Book, error) {
	return b.dataBase.SearchByAuthor(author, pagination, idx, limit)
}

func (b *bookServiceImpl) ByTitle(title string, pagination bool, idx, limit uint64) ([]models.Book, error) {
	return b.dataBase.SearchByAuthor(title, pagination, idx, limit)
}

func (b *bookServiceImpl) Update(book models.Book) {
	b.dataBase.Update(book.ID, book)
}

func NewBookService() BookService {
	bs := new(bookServiceImpl)
	bs.dataBase = DefaultDb

	return bs
}

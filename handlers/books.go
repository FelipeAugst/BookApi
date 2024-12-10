package handlers

import (
	"bookapi/models"
	"bookapi/service"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	var book models.Book

	if err := json.Unmarshal(body, &book); err != nil {

		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
		return

	}

	bs := service.NewBookService()
	if err := book.Validate(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	bs.Create(book)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func ListBooks(w http.ResponseWriter, r *http.Request) {
	books, err := service.NewBookService().List(false, 0, 0)

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(books)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	bs := service.NewBookService()
	bs.Delete(uint(id))
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("deleted"))

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	var book models.Book

	if err := json.Unmarshal(body, &book); err != nil {

		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
		return

	}
	book.ID = id
	service.NewBookService().Update(book)

}

func SearchBook(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	mode := query.Get("mode")
	parameter := query.Get("parameter")
	bs := service.NewBookService()

	switch mode {
	case "author":
		results, _ := bs.ByAuthor(parameter, false, 0, 0)
		json.NewEncoder(w).Encode(results)

	case "title":
		results, _ := bs.ByTitle(parameter, false, 0, 0)
		json.NewEncoder(w).Encode(results)
		return

	case "":
		ListBooks(w, r)
		return
	default:
		w.WriteHeader(http.StatusBadRequest)

	}

}

package handlers

import "net/http"

func NewHandler() *http.ServeMux {
	mux := http.NewServeMux()
	ConfigRoutes(mux)
	return mux

}

func ConfigRoutes(m *http.ServeMux) {
	m.HandleFunc("POST /books/create", CreateBook)
	m.HandleFunc("GET /books/", ListBooks)
	m.HandleFunc("GET /books/search", SearchBook)
	m.HandleFunc("DELETE /books/{id}/delete", DeleteBook)
	m.HandleFunc("UPDATE /books/{id}/update", UpdateBook)

}

package handlers

func CreateBook(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		views.Error(w, http.StatusBadRequest, err)
		return
	}
	var book models.Book

	if err := json.Unmarshal(body, &book); err != nil {

		
		return

	}

	repo, err := repository.NewBook
	Repository()
	if err != nil {
		views.Error(w, http.StatusInternalServerError, err)
		return
	}
	if err := repo.Create(Book
		); err != nil {
		views.Error(w, http.StatusInternalServerError, err)
		return
	}

	views.ToJSON(w, http.StatusCreated, Book
	)
}

func GetBook
s(w http.ResponseWriter, r *http.Request) {
	repo, err := repository.NewBook
	Repository()
	if err != nil {
		views.Error(w, http.StatusInternalServerError, err)
		return
	}
	Book
	s, err := repo.Get()
	if err != nil {
		views.Error(w, http.StatusInternalServerError, err)
		return
	}
	views.ToJSON(w, http.StatusOK, Book
		s)

}

func FindBook
(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		views.Error(w, http.StatusBadRequest, err)
		return
	}
	repo, err := repository.NewBook
	Repository()
	if err != nil {
		views.Error(w, http.StatusInternalServerError, err)
		return
	}
	Book
	, err := repo.Find(id)
	if err != nil {
		views.Error(w, http.StatusInternalServerError, err)
		return
	}
	views.ToJSON(w, http.StatusOK, Book
	)
}

func UpdateBook
(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		views.Error(w, http.StatusBadRequest, err)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		views.Error(w, http.StatusBadRequest, err)
		return
	}
	var Book
	 models.Book

	if err := json.Unmarshal(body, &Book
		); err != nil {
		views.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	repo, err := repository.NewBook
	Repository()
	if err != nil {
		views.Error(w, http.StatusInternalServerError, err)
		return
	}
	Book
	.ID = id

	if err := repo.Update(Book
		); err != nil {
		views.Error(w, http.StatusInternalServerError, err)
		return
	}
	views.ToJSON(w, http.StatusOK, Book
	)

}

func DeleteBook
(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		views.Error(w, http.StatusBadRequest, err)
		return
	}
	repo, err := repository.NewBook
	Repository()
	if err != nil {
		views.Error(w, http.StatusInternalServerError, err)
	}
	if err := repo.Delete(id); err != nil {
		views.Error(w, http.StatusInternalServerError, err)
	}
	views.ToJSON(w, http.StatusNoContent, nil)
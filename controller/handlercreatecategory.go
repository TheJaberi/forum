package forum

import (
	model "forum/model"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func HandlerCreateCategory(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/createcategory" {
		ErrorHandler(w, req, http.StatusNotFound)
		return
	}
	if req.Method != "POST" {
		ErrorHandler(w, req, http.StatusMethodNotAllowed)
		return
	}
	err := model.CreateCategory(req.FormValue("category"))
	if err != nil {
		ErrorHandler(w, req, http.StatusBadRequest)
		return
	}
	t, err := template.ParseFiles(HTMLs...)
	if err != nil {
		ErrorHandler(w, req, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "index.html", model.AllData)
}

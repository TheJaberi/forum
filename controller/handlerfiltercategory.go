package forum

import (
	model "forum/model"
	"html/template"
	"net/http"
)

func HandlerFilterCategory(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/filtercategory/" {
		ErrorHandler(w, req, http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		ErrorHandler(w, req, http.StatusMethodNotAllowed)
		return
	}
	t, err := template.ParseFiles(HTMLs...)
	if err != nil {
		ErrorHandler(w, req, http.StatusInternalServerError)
		return
	}
	err = model.FilterByCategory(req.FormValue("category"))
	if err != nil {
		ErrorHandler(w, req, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "index.html", model.AllData)
}

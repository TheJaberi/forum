package forum

import (
	"forum/functions"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
)
func HandlerCreateCategory(w http.ResponseWriter, req *http.Request){
	if req.URL.Path != "/createcategory" {
		ErrorHandler(w, req, http.StatusNotFound)
		return
	}
	if req.Method != "POST" {
		ErrorHandler(w, req, http.StatusMethodNotAllowed)
		return
	}
	t, err := template.ParseFiles(HTMLs...)
	if err != nil {
		ErrorHandler(w, req, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	categoryName := req.FormValue("category")
	forum.CreateCategory(categoryName)
	t.ExecuteTemplate(w, "main.html", forum.AllPosts)
}

package forum

import (
	"html/template"
	"net/http"
	forum"forum/functions"
)

func MainHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
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
	w.WriteHeader(http.StatusOK)
	title := req.FormValue("title")
	post := req.FormValue("post")
	forum.CreateTables(title, post)
	t.ExecuteTemplate(w, "main.html", nil)
}

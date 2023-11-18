package forum

import (
	"html/template"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/signup" {
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
	t.ExecuteTemplate(w, "register.html", nil)
}

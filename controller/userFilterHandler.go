package forum

import (
	model "forum/model"
	"html/template"
	"net/http"
)

func HandlerMyFilter(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/mylikes/" && req.URL.Path != "/myposts/" && req.URL.Path != "/mydislikes/" {
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
	err = model.FilterUserData(req.FormValue("user"), req.URL.Path) // depending on the url path that is displayed when each of the buttons is clicked the data well be filtered
	if err != nil {
		ErrorHandler(w, req, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "index.html", model.AllData)
}

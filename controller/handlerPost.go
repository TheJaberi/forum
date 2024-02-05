package forum

import (
	model "forum/model"
	"html/template"
	"net/http"
	"strconv"
)

func HandlerPost(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/post" {
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

	var postCategories []int
	for i := 1; i <= len(model.AllCategories); i++ {
		categorytmp := req.FormValue(strconv.Itoa(i)) // XXX is this correct form value?
		if categorytmp != "" {
			postCategories = append(postCategories, i)
		}
	}
	err = model.CreatePost(req.FormValue("title"), model.AdjustText(req.FormValue("post")), postCategories)
	if err != nil {
		ErrorHandler(w, req, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "index.html", model.AllData)
}

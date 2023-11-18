package forum

import (
	forum "forum/functions"
	"html/template"
	"net/http"
	"strconv"
)

func HandlerPost(w http.ResponseWriter, req *http.Request) {
	var postCategories []int
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
	w.WriteHeader(http.StatusOK)
	title := req.FormValue("title") // when the createpost button is clicked the title data is assigned to a variable
	body := req.FormValue("post")   // when the createpost button is clicked the body data is assigned to a variable
	for i := 1; i <= len(forum.AllCategories); i++ {
		categorytmp := req.FormValue(strconv.Itoa(i))
		if categorytmp != "" {
			postCategories = append(postCategories, i)
		}
	}
	forum.CreatePost(title, body, postCategories) // create post adds the title and body to the table in the database
	// MainHandler(w, req)
	t.ExecuteTemplate(w, "main.html", forum.AllData)
}

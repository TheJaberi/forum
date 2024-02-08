package forum

import (
	model "forum/model"
	"html/template"
	"net/http"
)

func HandlerComments(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/comment" {
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
	// when the createpost button is clicked the title data is assigned to a variable
	postData, err := model.CreateComment(req.FormValue("commentContent"), req.FormValue("postid")) // create post adds the title and body to the table in the database
	if err != nil {
		ErrorHandler(w, req, http.StatusBadRequest)
		return
	}
	err = model.GetPosts()
	if err != nil {
		ErrorHandler(w, req, http.StatusInternalServerError)
		return
	}
	postData.LoggedUser = true
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "postpage.html", postData)
}

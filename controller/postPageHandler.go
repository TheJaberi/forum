package forum

import (
	model "forum/model"
	"html/template"
	"net/http"
	"strconv"
)

func HandlerPostPage(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/postpage/" {
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

	postID, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil {
		ErrorHandler(w, req, http.StatusBadRequest)
		return
	}

	model.AllData.Postpage = model.AllPosts[postID-1]
	model.AllData.Postpage.LoggedUser = model.AllData.IsLogged

	if model.AllData.IsLogged {
		for i := 0; i < len(model.AllData.Postpage.Comments); i++ {
			model.AllData.Postpage.Comments[i].CommentLoggedUser = model.AllData.IsLogged
		}
	}
	w.WriteHeader(http.StatusOK)
	model.GetUserCommentInteractions()
	t.ExecuteTemplate(w, "postpage.html", model.AllData.Postpage)
}

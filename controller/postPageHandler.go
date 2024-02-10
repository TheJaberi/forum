package forum

import (
	// "fmt"
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
	err = model.GetPosts()
	if err != nil {
		ErrorHandler(w, req, http.StatusInternalServerError)
		return
	}
	postID, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil {
		ErrorHandler(w, req, http.StatusBadRequest)
		return
	}
	model.AllData.Postpage, err = model.GetPost((postID))
	if err != nil {
		ErrorHandler(w, req, http.StatusNotFound)
		return
	}
	model.AllData.Postpage.LoggedUser = model.AllData.IsLogged // TODO (Use checkCookie) if the user is registered the like and dislike buttons appear on the post's page
	err = model.GetPostComments(&model.AllData.Postpage)
	if err != nil {
		ErrorHandler(w, req, http.StatusNotFound)
		return
	}
	for i := 0; i < len(model.AllData.Postpage.Comments); i++ { // XXX Not sure what this does..
		model.AllData.Postpage.Comments[i].CommentLoggedUser = model.AllData.IsLogged
	}
	w.WriteHeader(http.StatusOK)
	model.GetUserPostInteractions(postID)
	t.ExecuteTemplate(w, "postpage.html", model.AllData.Postpage)
}

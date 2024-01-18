package forum

import (
	forum "forum/functions"
	"net/http"
)

func HandlerLogout(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/logout/" {
		ErrorHandler(w, req, http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		ErrorHandler(w, req, http.StatusMethodNotAllowed)
		return
	}
	// t, err := template.ParseFiles(HTMLs...)
	// if err != nil {
	// 	ErrorHandler(w, req, http.StatusInternalServerError)
	// 	return
	// }
	w.WriteHeader(http.StatusOK)
	var empty forum.User
	forum.AllData.LoggedUser = empty
	forum.AllData.IsLogged = false
	forum.LiveSession = forum.EmptySession
	http.Redirect(w, req, "/", http.StatusOK)
	forum.UpdatePosts()
	// t.ExecuteTemplate(w, "index.html", forum.AllData)
}

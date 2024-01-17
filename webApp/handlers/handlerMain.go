package forum

import (
	"fmt"
	forum "forum/functions"
	"html/template"
	"net/http"
)

func MainHandler(w http.ResponseWriter, req *http.Request) {
	if forum.AllData.IsLogged {
		cookie, err2 := req.Cookie("myCookies")
		if err2 != nil {
			fmt.Println(err2)
		}
		fmt.Println(cookie)
	}

	check := forum.CheckCookies(req)

	if check != nil {
		HandlerLogout(w, req)
	}
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
	forum.ViewCategory()
	forum.ViewPosts()
	forum.AllData.AllPosts = forum.AllPosts
	forum.AllData.AllCategories = forum.AllCategories
	// forum.AllData.LoggedUser = forum.LoggedUser
	// forum.AllData.LoggedUserID = forum.LoggedUser.Userid
	// forum.AllData.IsLogged = forum.LoggedUser.Registered
	t.ExecuteTemplate(w, "index.html", forum.AllData)
}

package forum

import (
	"fmt"
	forum "forum/functions"
	"html/template"
	"net/http"
)

func MainHandler(w http.ResponseWriter, req *http.Request) {
	if forum.AllData.IsLogged {
		_, err2 := req.Cookie("myCookies")
		if err2 != nil {
			forum.AllData.IsLogged = false
			forum.AllData.LoggedUser = forum.Empty
		}
	}
	if req.URL.Path != "/" {
		ErrorHandler(w, req, http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		ErrorHandler(w, req, http.StatusMethodNotAllowed)
		return
	}
	check := forum.CheckCookies(req)
	if check != nil {
		fmt.Println(check, "test")
		forum.AllData.IsLogged = false
		forum.AllData.LoggedUser = forum.Empty
	}
	t, err := template.ParseFiles(HTMLs...)
	if err != nil {
		ErrorHandler(w, req, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	forum.ViewCategory()
	forum.ViewPosts()
	forum.AllData.AllPosts = forum.RSort(forum.AllPosts)
	forum.AllData.AllCategories = forum.AllCategories
	forum.AllData.CategoryCheck = true
	// forum.AllData.LoggedUser = forum.LoggedUser
	// forum.AllData.LoggedUserID = forum.LoggedUser.Userid
	// forum.AllData.IsLogged = forum.LoggedUser.Registered
	sortby := req.FormValue("sortby")
	if sortby == "oldest" {
		forum.AllData.AllPosts = forum.AllPosts
	} else if sortby == "mostliked"{
		forum.AllData.AllPosts = forum.SortByLike(forum.AllPosts)
	} else if sortby == "mostdisliked"{
		forum.AllData.AllPosts = forum.SortByDislike(forum.AllPosts)
	}  else if sortby == "mostcommentedon"{
		forum.AllData.AllPosts = forum.SortByComment(forum.AllPosts)
	}
	if forum.LoginError2{
		forum.AllData.LoginError = true
	} else {
		forum.AllData.LoginError = false
	}
	t.ExecuteTemplate(w, "index.html", forum.AllData)
	forum.AllData.LoginError = false
	forum.LoginError2 = false
	forum.AllData.LoginErrorMsg = ""
}

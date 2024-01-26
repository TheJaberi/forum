package forum

import (
	forum "forum/functions"
	"html/template"
	"net/http"
	"strconv"
)

func MainHandler(w http.ResponseWriter, req *http.Request) {
	if forum.AllData.IsLogged {
		check := forum.CheckCookies(req)
		if check != nil {
			forum.AllData.LoggedUser = forum.Empty
			forum.AllData.IsLogged = false
			forum.LiveSession = forum.EmptySession
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

	t, err := template.ParseFiles(HTMLs...)
	if err != nil {
		ErrorHandler(w, req, http.StatusInternalServerError)
		return
	}
	forum.ViewCategory()
	forum.ViewPosts()
	forum.AllData.AllPosts = forum.RSort(forum.AllPosts)
	forum.AllData.AllCategories = forum.AllCategories
	// forum.AllData.LoggedUser = forum.LoggedUser
	// forum.AllData.LoggedUserID = forum.LoggedUser.Userid
	// forum.AllData.IsLogged = forum.LoggedUser.Registered
	sortby := req.FormValue("sortby")
	if sortby == "oldest" {
		forum.AllData.AllPosts = forum.AllPosts
	} else if sortby == "mostliked" {
		forum.AllData.AllPosts = forum.SortByLike(forum.AllPosts)
	} else if sortby == "mostdisliked" {
		forum.AllData.AllPosts = forum.SortByDislike(forum.AllPosts)
	} else if sortby == "mostcommentedon" {
		forum.AllData.AllPosts = forum.SortByComment(forum.AllPosts)
	}
	if forum.LoginError2 {
		forum.AllData.LoginError = true
	} else {
		forum.AllData.LoginError = false
	}
	category, _ := strconv.Atoi(req.FormValue("category")) // gets the data from the button clicked for filtering
	if category>0{
		var filteredPosts []forum.Post
	for i := 0; i < len(forum.AllPosts); i++ {
		for j := 0; j < len(forum.AllPosts[i].Category); j++ {
			if category == forum.AllPosts[i].Category[j].CategoryID { // loop over all the categories of all the posts if it matches "category" append the data of the post
				filteredPosts = append(filteredPosts, forum.AllPosts[i])
				break
			}
		}
	}
	forum.AllData.AllPosts = forum.RSort(filteredPosts)
}
	t.ExecuteTemplate(w, "index.html", forum.AllData)
	forum.AllData.LoginError = false
	forum.LoginError2 = false
	forum.AllData.LoginErrorMsg = ""
}

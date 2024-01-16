package forum

import (
	forum "forum/functions"
	"html/template"
	"net/http"
)

const PerPage = 10

func MainHandler(w http.ResponseWriter, req *http.Request) {
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
	forum.AllData.AllPosts = RSort(forum.AllPosts)
	forum.AllData.AllCategories = forum.AllCategories
	forum.AllData.CategoryCheck = true
	// forum.AllData.LoggedUser = forum.LoggedUser
	// forum.AllData.LoggedUserID = forum.LoggedUser.Userid
	// forum.AllData.IsLogged = forum.LoggedUser.Registered
	sortby := req.FormValue("sortby")
	if sortby == "oldest" {
		forum.AllData.AllPosts = forum.AllPosts
	} else if sortby == "mostliked"{
		forum.AllData.AllPosts = SortByLike(forum.AllPosts)
	} else if sortby == "mostdisliked"{
		forum.AllData.AllPosts = SortByDislike(forum.AllPosts)
	}  else if sortby == "mostcommentedon"{
		forum.AllData.AllPosts = SortByComment(forum.AllPosts)
	}
	t.ExecuteTemplate(w, "index.html", forum.AllData)
}

func RSort(list []forum.Post) []forum.Post {
	var arrAllPosts []forum.Post
	for i := len(list) - 1; i > 0; i-- {
		arrAllPosts = append(arrAllPosts, list[i])
	}
	return arrAllPosts
}

func SortByLike(list []forum.Post) []forum.Post {
	var arrAllPosts []forum.Post
	for i := 0; i >= 0; i++ {
		for j:=0;j<len(list);j++{
			if list[j].Likes == i {
				arrAllPosts = append(arrAllPosts, list[j])
			}
		}
		if len(arrAllPosts) >= len(list){
			break
		}
	}
	return RSort(arrAllPosts)
}
func SortByDislike(list []forum.Post) []forum.Post {
	var arrAllPosts []forum.Post
	for i := 0; i >= 0; i++ {
		for j:=0;j<len(list);j++{
			if list[j].Dislikes == i {
				arrAllPosts = append(arrAllPosts, list[j])
			}
		}
		if len(arrAllPosts) >= len(list){
			break
		}
	}
	return RSort(arrAllPosts)
}
func SortByComment(list []forum.Post) []forum.Post {
	var arrAllPosts []forum.Post
	for i := 0; i >= 0; i++ {
		for j:=0;j<len(list);j++{
			if len(list[j].Comments) == i {
				arrAllPosts = append(arrAllPosts, list[j])
			}
		}
		if len(arrAllPosts) >= len(list){
			break
		}
	}
	return RSort(arrAllPosts)
}
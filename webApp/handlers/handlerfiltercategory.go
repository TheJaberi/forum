package forum

import (
	forum "forum/functions"
	"html/template"
	"net/http"
	"strconv"
)

func HandlerFilterCategory(w http.ResponseWriter, req *http.Request) {
	var filteredPosts []forum.Post
	if req.URL.Path != "/filtercategory/"{
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
		// var filterCategory []forum.Category
	category, _ := strconv.Atoi(req.FormValue("category")) // gets the data from the button clicked for filtering
	for i := 0; i < len(forum.AllPosts); i++ {
		for j := 0; j < len(forum.AllPosts[i].Category); j++ {
			if category == forum.AllPosts[i].Category[j].CategoryID { // loop over all the categories of all the posts if it matches "category" append the data of the post
				filteredPosts = append(filteredPosts, forum.AllPosts[i])
				// if filterCategory == nil {
				// filterCategory = append(filterCategory, forum.AllPosts[i].Category[j])}
				break
			}
		}
	}
	w.WriteHeader(http.StatusOK)
	forum.AllData.AllPosts = forum.RSort(filteredPosts)
	// forum.AllData.AllCategories = filterCategory
	forum.AllData.CategoryCheck = false
	forum.AllData.LoggedUser = forum.LoggedUser
	t.ExecuteTemplate(w, "index.html", forum.AllData) // execute the main html with only the filtered posts
}

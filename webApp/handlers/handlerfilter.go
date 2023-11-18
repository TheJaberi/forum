package forum

import (
	// "fmt"
	"fmt"
	forum "forum/functions"
	"html/template"
	"net/http"
	"strconv"
)

func HandlerFilter(w http.ResponseWriter, req *http.Request){
var filteredPosts []forum.Post
	if req.URL.Path != "/filter/" {
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
test := req.FormValue("category")
fmt.Println(test)
category, _ := strconv.Atoi(req.FormValue("category"))
fmt.Println(category)
for i:=0;i<len(forum.AllPosts);i++{
	for j:=0;j<len(forum.AllPosts[i].Category);j++{
	if category == forum.AllPosts[i].Category[j].CategoryID {
		filteredPosts = append(filteredPosts, forum.AllPosts[i])
		break
	}
}
}
forum.AllData.AllPosts = filteredPosts
forum.AllData.AllCategories = forum.AllCategories
t.ExecuteTemplate(w, "main.html", forum.AllData)
}
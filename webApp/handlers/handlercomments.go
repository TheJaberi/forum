package forum

import (
	"fmt"
	forum "forum/functions"
	"html/template"
	"net/http"
	"strconv"
)

func HandlerComments(w http.ResponseWriter, req *http.Request) {
	fmt.Println("11")
	if req.URL.Path != "/comment/" {
		fmt.Println("11")
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
	post_id, _ := strconv.Atoi(req.FormValue("postid")) // when the createpost button is clicked the title data is assigned to a variable
	commentContent := req.FormValue("commentContent")
	fmt.Println(post_id, commentContent)
	forum.CreateComment(commentContent, post_id) // create post adds the title and body to the table in the database
	postData := forum.AllPosts[post_id]
	t.ExecuteTemplate(w, "postpage.html", postData)
}

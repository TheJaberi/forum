package forum

import (
	"fmt"
	forum "forum/functions"
	"html/template"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func HandlerLikes(w http.ResponseWriter, req *http.Request){
	if req.URL.Path != "/dislike/" && req.URL.Path != "/like/" {
		fmt.Println("@222")
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
	addPost_id, _ := strconv.Atoi(req.FormValue("postInteraction"))
	remPost_id, _ := strconv.Atoi(req.FormValue("removeInteraction"))
	user_id := forum.LoggedUser.Userid
	if addPost_id > remPost_id{
		if req.URL.Path == "/like/"{
		 if !forum.AllPosts[addPost_id-1].UserDislike{
		forum.InsertInteraction(addPost_id, user_id, 1)
		forum.AllPosts[addPost_id-1].Userlike =  true
	} else {
		forum.UpdateInteraction(addPost_id, user_id, 1) // function not created
		forum.AllPosts[addPost_id-1].Userlike =  true
		forum.AllPosts[addPost_id-1].UserDislike =  false
	}
} else {
	if !forum.AllPosts[addPost_id-1].Userlike{
		forum.InsertInteraction(addPost_id, user_id, 0)
		forum.AllPosts[addPost_id-1].UserDislike =  true
	} else {
		forum.UpdateInteraction(addPost_id, user_id, 0)
		forum.AllPosts[addPost_id-1].UserDislike =  true
		forum.AllPosts[addPost_id-1].Userlike =  false
	}
}
} else{
	fmt.Println("test")
		forum.RemoveInteraction(remPost_id, user_id)
}
	t.ExecuteTemplate(w, "postpage.html", forum.AllData)
}
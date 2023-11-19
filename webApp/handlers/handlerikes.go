package forum

import (
	forum "forum/functions"
	_"github.com/mattn/go-sqlite3"
	"strconv"
	"html/template"
	"net/http"
)

func HandlerLikes(w http.ResponseWriter, req *http.Request){
	if req.URL.Path != "/dislike" && req.URL.Path != "/like" {
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
		if req.URL.Path != "/like"{
		 if !forum.AllPosts[addPost_id].UserDislike{
		forum.InsertInteraction(addPost_id, user_id, true)
		forum.AllPosts[addPost_id].Userlike =  true
		forum.AllPosts[addPost_id].Likes++ 
	} else {
		forum.UpdateInteraction(addPost_id, user_id, true) // function not created
		forum.AllPosts[addPost_id].Userlike =  true
		forum.AllPosts[addPost_id].UserDislike =  false
		forum.AllPosts[addPost_id].Likes++ 
		forum.AllPosts[addPost_id].Dislikes-- 
	}
} else {
	if !forum.AllPosts[addPost_id].Userlike{
		forum.InsertInteraction(addPost_id, user_id, false)
		forum.AllPosts[addPost_id].UserDislike =  true
		forum.AllPosts[addPost_id].Dislikes++ 
	} else {
		forum.UpdateInteraction(addPost_id, user_id, false)
		forum.AllPosts[addPost_id].UserDislike =  true
		forum.AllPosts[addPost_id].Userlike =  false
		forum.AllPosts[addPost_id].Dislikes++
		forum.AllPosts[addPost_id].Likes--
	}
}
} else{
		forum.RemoveInteraction(addPost_id, user_id)
}
		forum.AllPosts[addPost_id].Dislikes++
	t.ExecuteTemplate(w, "postpage.html", forum.AllData)
}
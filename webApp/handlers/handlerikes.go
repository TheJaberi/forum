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
	addPost_id--
	remPost_id--
	fmt.Println(addPost_id)
	fmt.Println(remPost_id)
	fmt.Println(forum.AllPosts)
	user_id := forum.LoggedUser.Userid
	if addPost_id > remPost_id{
		if req.URL.Path == "/like"{
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
		forum.RemoveInteraction(remPost_id, user_id)
}
	t.ExecuteTemplate(w, "postpage.html", forum.AllData)
}
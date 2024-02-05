package forum
import (
	// "fmt"
	forum "forum/functions"
	"html/template"
	"net/http"
	"strconv"
	_ "github.com/mattn/go-sqlite3"
)
func HandlerLikes(w http.ResponseWriter, req *http.Request){
	var postData forum.Post
	if !forum.AllData.IsLogged{
		ErrorHandler(w, req, http.StatusNotFound)
		return
	}
	if req.URL.Path != "/dislike/" && req.URL.Path != "/like/" {
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
	addPost_id, _ := strconv.Atoi(req.FormValue("postInteraction")) // post interaction handles the data from like or dislike button if the user logged hasn't already clicked on it 
	remPost_id, _ := strconv.Atoi(req.FormValue("removeInteraction")) // remove interaction handles the data from like or dislike button if the user logged has already clicked on it 
	user_id := forum.LoggedUser.Userid
	if addPost_id > remPost_id{ // which ever value is greater determines whether to add or remove 
		postData = forum.AllPosts[addPost_id-1]
		if req.URL.Path == "/like/"{
		 if !forum.AllPosts[addPost_id-1].UserDislike{
		forum.InsertInteraction(addPost_id, user_id, 1, true, 0) // insert adds the interaction to the database 1 is like 0 is dislike
		forum.AllPosts[addPost_id-1].Userlike =  true // changes the post like or dislike for the logged in user in the all posts var
	} else {
		forum.UpdateInteraction(addPost_id, user_id, 1, true, 0) // update is used if a like has to be changed to a dislike or vice versa
		forum.AllPosts[addPost_id-1].Userlike =  true
		forum.AllPosts[addPost_id-1].UserDislike =  false
	}
} else {
	if !forum.AllPosts[addPost_id-1].Userlike{
		forum.InsertInteraction(addPost_id, user_id, 0, true, 0)
		forum.AllPosts[addPost_id-1].UserDislike =  true
	} else {
		forum.UpdateInteraction(addPost_id, user_id, 0, true, 0)
		forum.AllPosts[addPost_id-1].UserDislike =  true
		forum.AllPosts[addPost_id-1].Userlike =  false
	}
}
postData = forum.AllPosts[addPost_id-1]
} else{
		forum.RemoveInteraction(remPost_id, user_id, true, 0) //remove is greater means there is already an interaction that needs to be removed
		forum.AllPosts[remPost_id-1].Userlike =  false
		forum.AllPosts[remPost_id-1].UserDislike =  false
		postData = forum.AllPosts[remPost_id-1]
}
	postData.LoggedUser = true
	t.ExecuteTemplate(w, "postpage.html", postData)
}
package main

import (
	"fmt"
	"log"
	"net/http"

	forumfunc "forum/functions"
	forum "forum/webApp/handlers"
)

// Whatever needs to load before the server starts (Files/APIs)
func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	forum.StaticFileLoader()
	forumfunc.DatabaseLoader()
}

func main() {
	const port = ":8080"
	http.HandleFunc("/", forum.MainHandler)
	fmt.Println("http://localhost" + port)
	// forumfunc.CreateTables() // create table creates the database and the tables for the project
	http.HandleFunc("/createcategory", forum.HandlerCreateCategory)  // for the admin only to create new categories
	http.HandleFunc("/postpage/", forum.HandlerPostPage)             // handles the post that is clicked on in the homepage
	http.HandleFunc("/filtercategory/", forum.HandlerFilterCategory) // handles the filtering by category
	http.HandleFunc("/myposts/", forum.HandlerMyFilter)              // handles the filtering by user's posts, likes or dislikes
	http.HandleFunc("/mylikes/", forum.HandlerMyFilter)
	http.HandleFunc("/logout/", forum.HandlerLogout)
	http.HandleFunc("/comment", forum.HandlerComments)
	http.HandleFunc("/mydislikes/", forum.HandlerMyFilter)
	http.HandleFunc("/like/", forum.HandlerLikes) // handles the function which adds the interaction to the database
	http.HandleFunc("/dislike/", forum.HandlerLikes)
	http.HandleFunc("/commentlike/", forum.HandlerCommentsLikes)
	http.HandleFunc("/commentdislike/", forum.HandlerCommentsLikes)
	http.HandleFunc("/register", forum.HandlerRegister) // HandlerRegister has function NewUser which adds the data for the user to the database
	http.HandleFunc("/login", forum.HandlerLogin)       // HandlerLogin checks if the user is registered, if so it adds his data to a Global variable
	http.HandleFunc("/post", forum.HandlerPost)         // HandlerPost adds the data in the post to the database
	log.Fatal(http.ListenAndServe(port, nil))
}

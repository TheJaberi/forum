package main

import (
	"fmt"
	"log"
	"net/http"

	controller "forum/controller"
	model "forum/model"
)

// Whatever needs to load before the server starts (Files/APIs)
func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	controller.StaticFileLoader()
	model.DatabaseLoader()
}

func main() {
	const port = ":8080"
	http.HandleFunc("/", controller.MainHandler)
	fmt.Println("http://localhost" + port)
	// model.CreateTables() // create table creates the database and the tables for the project
	http.HandleFunc("/createcategory", controller.HandlerCreateCategory)  // for the admin only to create new categories
	http.HandleFunc("/postpage/", controller.HandlerPostPage)             // handles the post that is clicked on in the homepage
	http.HandleFunc("/filtercategory/", controller.HandlerFilterCategory) // handles the filtering by category
	http.HandleFunc("/myposts/", controller.HandlerMyFilter)              // handles the filtering by user's posts, likes or dislikes
	http.HandleFunc("/mylikes/", controller.HandlerMyFilter)
	http.HandleFunc("/logout/", controller.HandlerLogout)
	http.HandleFunc("/comment", controller.HandlerComments)
	http.HandleFunc("/mydislikes/", controller.HandlerMyFilter)
	http.HandleFunc("/like/", controller.HandlerLikes) // handles the function which adds the interaction to the database
	http.HandleFunc("/dislike/", controller.HandlerLikes)
	http.HandleFunc("/commentlike/", controller.HandlerCommentsLikes)
	http.HandleFunc("/commentdislike/", controller.HandlerCommentsLikes)
	http.HandleFunc("/register", controller.HandlerRegister) // HandlerRegister has function NewUser which adds the data for the user to the database
	http.HandleFunc("/login", controller.HandlerLogin)       // HandlerLogin checks if the user is registered, if so it adds his data to a Global variable
	http.HandleFunc("/post", controller.HandlerPost)         // HandlerPost adds the data in the post to the database
	log.Fatal(http.ListenAndServe(port, nil))
}

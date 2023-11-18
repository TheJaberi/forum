package main

import (
	"fmt"
	forum "forum/webApp/handlers"
	forumfunc "forum/functions"
	"log"
	"net/http"
)

// Whatever needs to load before the server starts (Files/APIs)
func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	forum.StaticFileLoader()
}

func main() {
	const port = ":8080"
	http.HandleFunc("/", forum.MainHandler) // MainHandler executes main.html an has the function create table which creates the database
	fmt.Println("http://localhost" + port)
	forumfunc.CreateTables() // create table creates the database and the tables for the project
	http.HandleFunc("/createcategory", forum.HandlerCreateCategory)
	http.HandleFunc("/postpage/", forum.HandlerPostPage)
	http.HandleFunc("/filter/", forum.HandlerFilter)
	http.HandleFunc("/register", forum.HandlerRegister) // HandlerRegister has function NewUser which adds the data for the user to the database
	http.HandleFunc("/login", forum.HandlerLogin) // HandlerLogin checks if the user is registered, if so it adds his data to a Global variable
	http.HandleFunc("/post", forum.HandlerPost)  // HandlerPost adds the data in the post to the database
	log.Fatal(http.ListenAndServe(port, nil))
}

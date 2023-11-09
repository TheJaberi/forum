package main

import (
	"fmt"
	forum "forum/webApp/handlers"
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
	http.HandleFunc("/", forum.MainHandler)
	fmt.Println("http://localhost" + port)
	http.HandleFunc("/post", forum.Posthandler)
	log.Fatal(http.ListenAndServe(port, nil))
}

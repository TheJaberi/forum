package main

import (
	"fmt"
	db "forum/database"
	webApp "forum/webApp/handlers"
	"log"
	"net/http"
)

// Whatever needs to load before the server starts (Files/APIs)
func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	webApp.StaticFileLoader()
	db.DatabaseLoader()
}

func main() {
	const port = ":8080"
	http.HandleFunc("/", webApp.PathHandler)
	fmt.Println("http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}

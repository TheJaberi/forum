package main

import (
	"fmt"
	"log"
	"net/http"
	"crypto/tls"
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
	const port = ":443"
	http.HandleFunc("/", forum.MainHandler)
	fmt.Println("http://localhost" + port)
	// forumfunc.CreateTables() // create table creates the database and the tables for the project
	http.HandleFunc("/createcategory", forum.HandlerCreateCategory)  // for the admin only to create new categories
	http.HandleFunc("/postpage/", forum.HandlerPostPage)             // handles the post that is clicked on in the homepage
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
	mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
        w.Write([]byte("This is an example server.\n"))
    })
    cfg := &tls.Config{
        MinVersion:               tls.VersionTLS12,
        CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
        PreferServerCipherSuites: true,
        CipherSuites: []uint16{
            tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
            tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
            tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
            tls.TLS_RSA_WITH_AES_256_CBC_SHA,
        },
    }
    srv := &http.Server{
        Addr:         ":443",
        Handler:      mux,
        TLSConfig:    cfg,
        TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
    }
    log.Fatal(srv.ListenAndServeTLS("tls.crt", "tls.key"))
	err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
	if err != nil {
        fmt.Println("ListenAndServe: ", err)
	}
}

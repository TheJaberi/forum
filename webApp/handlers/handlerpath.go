package forum

import (
	db "forum/database"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func PathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/": // home page
		MainHandler(w, r)
	case "/sign_up": // sign up page
		RegisterHandler(w, r, db.DB)
	case "/sign_in": // sign in page
		LoginHandler(w, r, db.DB)
		/*
			case "/create_post": // welcome page
				WelcomeHundler(w, r)
			case "/create_comment": // filter page
				FilterHundler(w, r, data)
			case "/post": // search page
				SearchHundler(w, r, data)
			case "/400":
				ErrorHandler(w, r, 400)
			case "/500":
				ErrorHandler(w, r, 500)
		*/
	default: // any invalid path will be redirected to 404 error handler
		ErrorHandler(w, r, http.StatusNotFound)
	}
}

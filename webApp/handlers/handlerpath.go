package forum

import "net/http"

func PathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/": // home page
		MainHandler(w, r)
	// case "/sign_up": // artist page
	// 	ArtistHandler(w, r, data)
	// case "/sign_in": // about page
	// 	AboutHandler(w, r)
	// case "/create_post": // welcome page
	// 	WelcomeHundler(w, r)
	// case "/create_comment": // filter page
	// 	FilterHundler(w, r, data)
	// case "/post": // search page
	// 	SearchHundler(w, r, data)
	case "/400":
		ErrorHandler(w, r, 400)
	case "/500":
		ErrorHandler(w, r, 500)
	default: // any invalid path will be redirected to 404 error handler
		ErrorHandler(w, r, 404)
	}
}

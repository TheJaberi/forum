package forum

import (
	"net/http"
)

func PathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		MainHandler(w, r)
	case "/signup":
		RegisterHandler(w, r)
	case "/400":
		ErrorHandler(w, r, 400)
	case "/500":
		ErrorHandler(w, r, 500)
	default:
		ErrorHandler(w, r, 404)
	}
}

package forum

import "net/http"

var HTMLs []string

func StaticFileLoader() {
	if HTMLs == nil {
		HTMLs = []string{
			// Add new html / template names here
			"../webApp/static/main.html",
			"../webApp/static/error.html",
			"../webApp/static/postpage.html",
		}
	}
	cssFiles := http.FileServer(http.Dir("../webApp/static/css"))
	http.Handle("/css/", http.StripPrefix("/css/", cssFiles))
}

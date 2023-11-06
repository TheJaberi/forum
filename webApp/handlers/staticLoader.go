package forum

import "net/http"

var HTMLs []string

func StaticFileLoader() {
	if HTMLs == nil {
		HTMLs = []string{
			// Add new html / template names here
			"../webapp/static/main.html",
		}
	}
	cssFiles := http.FileServer(http.Dir("../webapp/static/css"))
	http.Handle("/css/", http.StripPrefix("/css/", cssFiles))
}

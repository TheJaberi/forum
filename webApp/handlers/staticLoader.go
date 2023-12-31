package forum

import "net/http"

var HTMLs []string

func StaticFileLoader() {
	if HTMLs == nil {
		HTMLs = []string{
			// Add new html / template names here
			"../webApp/static/index.html",
			"../webApp/static/main.html",
			"../webApp/static/error2.html",
			"../webApp/static/postpage.html",
			"../webApp/static/register2.html",
			"../webApp/static/subscribe2.html",
		}
	}
	cssFiles := http.FileServer(http.Dir("../webApp/static/css"))
	jsFiles := http.FileServer(http.Dir("../webApp/static/js"))
	imgFiles := http.FileServer(http.Dir("../webApp/static/img"))
	fontFiles := http.FileServer(http.Dir("../webApp/static/fonts"))
	http.Handle("/css/", http.StripPrefix("/css/", cssFiles))
	http.Handle("/js/", http.StripPrefix("/js/", jsFiles))
	http.Handle("/img/", http.StripPrefix("/img/", imgFiles))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", fontFiles))
}

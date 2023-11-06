package forum

import (
	"html/template"
	"log"
	"net/http"
)

func MainHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		log.Fatalf("Wrong URL, 404")
	}
	if req.Method != "GET" {
		log.Fatalf("Wrong Method, 405")
	}
	t, err := template.ParseFiles(HTMLs...)
	if err != nil {
		log.Fatalf("Files Not Parsed, 505")
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "main.html", nil)
}

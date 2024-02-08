package forum

import (
	"html/template"
	"log"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request, statusError int) {
	var errResponse struct {
		StatusCode    int
		StatusMessage string
	}
	errResponse.StatusCode = statusError
	switch {
	case statusError == http.StatusBadRequest:
		errResponse.StatusMessage = "Bad User Request"
	case statusError == http.StatusUnauthorized:
		errResponse.StatusMessage = "Unauthorized"
	case statusError == http.StatusForbidden:
		errResponse.StatusMessage = "Forbidden"
	case statusError == http.StatusNotFound:
		errResponse.StatusMessage = "Not Found"
	case statusError == http.StatusMethodNotAllowed:
		errResponse.StatusMessage = "Method Not Allowed"
	case statusError == http.StatusInternalServerError:
		errResponse.StatusMessage = "Internal Server Error"
	default:
		errResponse.StatusMessage = "Unknown Error, Contact The Pope for Answers"
	}
	t, err := template.ParseFiles(HTMLs...)
	if err != nil {
		log.Fatalf("Files Not Parsed, 505")
	}
	t.ExecuteTemplate(w, "error.html", errResponse)
}

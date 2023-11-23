package forum

import (
	"database/sql"
	"forum"
	UserDb "forum/database"
	"html/template"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	err := r.ParseForm()
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	data := forum.Login{
		Email:    r.PostFormValue("user_email2"),
		Password: r.PostFormValue("password2"),
	}
	if UserDb.UserDbLogin(data, db) != nil {
		ErrorHandler(w, r, http.StatusNotAcceptable)
		return
	}

	t, err := template.ParseFiles(HTMLs...)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "test.html", nil)
}

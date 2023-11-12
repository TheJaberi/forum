package forum

import (
	"database/sql"
	"html/template"
	"net/http"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user_name := r.PostFormValue("user_name")
	user_email := r.PostFormValue("user_email")
	user_pass := r.PostFormValue("password")
	user_type := "member"
	println(user_name, user_email, user_pass, user_type)
	sqlStmt, err := db.Prepare("INSERT INTO users (user_name, user_email, user_pass, user_type) VALUES (?, ?, ?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sqlStmt.Exec(user_name, user_email, user_pass, user_type)
	t, err := template.ParseFiles(HTMLs...)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "test.html", nil)
}
